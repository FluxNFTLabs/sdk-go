package main

import (
	"context"
	"cosmossdk.io/math"
	"encoding/hex"
	"encoding/json"
	"fmt"
	astromeshtypes "github.com/FluxNFTLabs/sdk-go/chain/modules/astromesh/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	"math/big"
	"os"
	"sort"
	"strings"

	evmtypes "github.com/FluxNFTLabs/sdk-go/chain/modules/evm/types"
	chaintypes "github.com/FluxNFTLabs/sdk-go/chain/types"
	"github.com/FluxNFTLabs/sdk-go/client/common"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	chainclient "github.com/FluxNFTLabs/sdk-go/client/chain"
)

type PairKey struct {
	Currency0   ethcommon.Address // base denom addr
	Currency1   ethcommon.Address // quote denom addr
	Fee         *big.Int          // 3000 -> 0.3%
	TickSpacing *big.Int
	Hooks       ethcommon.Address // hook contract addr
}

func main() {
	network := common.LoadNetwork("local", "")
	kr, err := keyring.New(
		"fluxd",
		"file",
		os.Getenv("HOME")+"/.fluxd",
		strings.NewReader("12345678\n"),
		chainclient.GetCryptoCodec(),
	)
	if err != nil {
		panic(err)
	}

	// init grpc connection
	cc, err := grpc.Dial("localhost:9900", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	// init client ctx
	clientCtx, senderAddress, err := chaintypes.NewClientContext(
		network.ChainId,
		"user1",
		kr,
	)
	if err != nil {
		panic(err)
	}
	clientCtx = clientCtx.WithGRPCClient(cc)

	// init chain client
	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		common.OptionGasPrices("500000000lux"),
	)
	if err != nil {
		panic(err)
	}

	// init evm client
	evmClient := evmtypes.NewQueryClient(cc)
	astromeshClient := astromeshtypes.NewQueryClient(cc)

	// read bytecode
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	bz, err := os.ReadFile(dir + "/examples/chain/28_DeployUniswap/PoolManager.json")
	if err != nil {
		panic(err)
	}

	var compData map[string]interface{}
	err = json.Unmarshal(bz, &compData)
	abiBz, err := json.Marshal(compData["abi"].([]interface{}))
	if err != nil {
		panic(err)
	}
	contractABI, err := abi.JSON(strings.NewReader(string(abiBz)))
	if err != nil {
		panic(err)
	}

	// parse contract addr
	PoolManagerContractAddr := "07aa076883658b7ed99d25b1e6685808372c8fe2"
	PoolManagerContractAddrBz, err := hex.DecodeString(PoolManagerContractAddr)
	if err != nil {
		panic(err)
	}

	// query tick spacing
	calldata, err := contractABI.Pack("MAX_TICK_SPACING")
	if err != nil {
		panic(err)
	}
	queryRes, err := evmClient.ContractQuery(context.Background(), &evmtypes.ContractQueryRequest{
		Address:  PoolManagerContractAddr,
		Calldata: calldata,
	})
	if err != nil {
		panic(err)
	}
	queryOutput, err := contractABI.Unpack("MAX_TICK_SPACING", queryRes.Output)
	if err != nil {
		panic(err)
	}
	fmt.Println("query MAX_TICK_SPACING:", queryOutput)

	calldata, err = contractABI.Pack("MIN_TICK_SPACING")
	if err != nil {
		panic(err)
	}
	queryRes, err = evmClient.ContractQuery(context.Background(), &evmtypes.ContractQueryRequest{
		Address:  PoolManagerContractAddr,
		Calldata: calldata,
	})
	if err != nil {
		panic(err)
	}
	queryOutput, err = contractABI.Unpack("MIN_TICK_SPACING", queryRes.Output)
	if err != nil {
		panic(err)
	}
	fmt.Println("query MIN_TICK_SPACING:", queryOutput)

	//perform astrotransfer to evm planes for cosmos tokens
	tokens := []string{"btc", "eth", "sol", "usdt"}
	tokenContracts := map[string]string{}
	for _, token := range tokens {
		amount, _ := math.NewIntFromString("100000000000000000000") // 100 * 10^18
		_, err := chainClient.SyncBroadcastMsg(&astromeshtypes.MsgAstroTransfer{
			Sender:   senderAddress.String(),
			Receiver: senderAddress.String(),
			SrcPlane: astromeshtypes.Plane_COSMOS,
			DstPlane: astromeshtypes.Plane_EVM,
			Coin: sdk.Coin{
				Denom:  token,
				Amount: amount,
			},
		})
		if err != nil {
			panic(err)
		}
		denomLink, err := astromeshClient.DenomLink(context.Background(), &astromeshtypes.QueryDenomLinkRequest{
			SrcPlane: astromeshtypes.Plane_COSMOS,
			DstPlane: astromeshtypes.Plane_EVM,
			SrcAddr:  token,
		})
		if err != nil {
			panic(err)
		}
		tokenContracts[token] = denomLink.DstAddr
		fmt.Println(fmt.Sprintf("%s on evm: %s, decimals %d", token, denomLink.DstAddr, denomLink.DstDecimals))
	}

	// initialize pairs
	prices := map[string]uint64{
		"btc": 69000,
		"eth": 3900,
		"sol": 185,
	}
	for denom, contractAddr := range tokenContracts {
		if denom == "usdt" {
			continue
		}
		// uniswap make sure only a single pair of 2 addresses can exist by comparing token addresses
		usdtAddr := tokenContracts["usdt"]
		currencies := []string{usdtAddr, contractAddr}
		sort.Strings(currencies)
		pairKey := &PairKey{
			Currency0:   ethcommon.Address(ethcommon.Hex2Bytes(currencies[0])),
			Currency1:   ethcommon.Address(ethcommon.Hex2Bytes(currencies[1])),
			Fee:         big.NewInt(3000),
			TickSpacing: big.NewInt(60),
			Hooks:       ethcommon.Address(make([]byte, 20)),
		}

		// 1 btc = 69000 usdt
		sqrtPriceX96Int := computeSqrtPriceX96Int(prices[denom])
		hookData := []byte{}
		calldata, err = contractABI.Pack("initialize", pairKey, sqrtPriceX96Int, hookData)
		if err != nil {
			panic(err)
		}

		msg := &evmtypes.MsgExecuteContract{
			Sender:          senderAddress.String(),
			ContractAddress: PoolManagerContractAddrBz,
			Calldata:        calldata,
		}

		res, err := chainClient.SyncBroadcastMsg(msg)
		if err != nil {
			panic(err)
		}

		// unpack execute return data
		var msgData sdk.TxMsgData
		var executeRes evmtypes.MsgExecuteContractResponse
		resBz := ethcommon.Hex2Bytes(res.TxResponse.Data)
		proto.Unmarshal(resBz, &msgData)
		proto.Unmarshal(msgData.MsgResponses[0].Value, &executeRes)
		initializeRes, err := contractABI.Unpack("initialize", executeRes.Output)
		if err != nil {
			panic(err)
		}
		tickId := initializeRes[0].(*big.Int)
		fmt.Println(fmt.Sprintf("pair %s-%s initialized, tick size %s: %s", denom, "usdt", tickId, res.TxResponse.TxHash))
	}

	// provide liquidity to pairs
}

func computeSqrtPriceX96Int(p uint64) *big.Int {
	sqrtPrice := new(big.Float).Sqrt(big.NewFloat(float64(p)))
	factor := new(big.Float).SetInt(new(big.Int).Lsh(big.NewInt(1), 96))
	sqrtPrice.Mul(sqrtPrice, factor)
	sqrtPriceX96Int := new(big.Int)
	price, _ := sqrtPrice.Int(sqrtPriceX96Int)
	return price
}