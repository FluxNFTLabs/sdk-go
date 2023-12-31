package types

import (
	"cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreate{}
var _ sdk.Msg = &MsgPurchaseShares{}
var _ sdk.Msg = &MsgTransferShares{}
var _ sdk.Msg = &MsgDepositShares{}
var _ sdk.Msg = &MsgWithdrawShares{}
var _ sdk.Msg = &MsgSponsor{}

// ValidateBasic implements the Msg.ValidateBasic method.
func (m MsgCreate) ValidateBasic() error {
	if len(m.ClassId) == 0 {
		return ErrEmptyClassID
	}

	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", m.Sender)
	}

	if !m.Supply.GT(sdkmath.ZeroInt()) {
		return errors.Wrapf(sdkerrors.ErrInvalidCoins, "Invalid nft supply (%s)", m.Supply)
	}

	minOwnerEquityPercent := sdkmath.NewIntFromUint64(33)
	if m.OwnerEquityPercent.LT(minOwnerEquityPercent) {
		return errors.Wrapf(ErrInvalidOwnerEquity, "Minimum owner equity percent is %s, got %s", minOwnerEquityPercent.String(), m.OwnerEquityPercent.String())
	}

	if !m.InitialPrice.Amount.GT(sdkmath.ZeroInt()) {
		return errors.Wrapf(sdkerrors.ErrInvalidCoins, "Invalid nft initial price (%s)", m.Supply)
	}

	if m.ISOTimestamp == 0 {
		return errors.Wrapf(sdkerrors.ErrNotSupported, "Invalid ISO timestamp (%d)", m.ISOTimestamp)
	}

	minSuccessPercent := uint64(66)
	if m.ISOSuccessPercent < minSuccessPercent {
		return errors.Wrapf(ErrInvalidISO, "ISO success threshold cannot be smaller than minimum percent (%d)", minSuccessPercent)
	}

	err = sdk.ValidateDenom(m.AcceptedPaymentDenom)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidCoins, "Invalid denom format (%s)", m.AcceptedPaymentDenom)
	}

	minDividendInterval := uint64(10)
	if m.DividendInterval < minDividendInterval {
		return errors.Wrapf(sdkerrors.ErrNotSupported, "Invalid dividend interval (%d), minimal interval is (%d)", m.DividendInterval, minDividendInterval)
	}

	return nil
}

// GetSigners returns the expected signers for MsgSend.
func (m MsgCreate) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{signer}
}

// ValidateBasic implements the Msg.ValidateBasic method.
func (m MsgPurchaseShares) ValidateBasic() error {
	if len(m.ClassId) == 0 {
		return ErrEmptyClassID
	}

	if len(m.Id) == 0 {
		return ErrEmptyNFTID
	}

	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", m.Sender)
	}

	if !m.Shares.GT(sdkmath.ZeroInt()) {
		return errors.Wrapf(sdkerrors.ErrInvalidCoins, "Invalid share amount (%s)", m.Shares)
	}

	return nil
}

// GetSigners returns the expected signers for MsgSend.
func (m MsgPurchaseShares) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{signer}
}

// ValidateBasic implements the Msg.ValidateBasic method.
func (m MsgTransferShares) ValidateBasic() error {
	if len(m.ClassId) == 0 {
		return ErrEmptyClassID
	}

	if len(m.Id) == 0 {
		return ErrEmptyNFTID
	}

	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", m.Sender)
	}

	_, err = sdk.AccAddressFromBech32(m.Receiver)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid receiver address (%s)", m.Receiver)
	}

	if !m.Shares.GT(sdkmath.ZeroInt()) {
		return errors.Wrapf(sdkerrors.ErrInvalidCoins, "Invalid share amount (%s)", m.Shares)
	}

	return nil
}

// GetSigners returns the expected signers for MsgSend.
func (m MsgTransferShares) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{signer}
}

// MsgDepositShares implements the Msg.ValidateBasic method.
func (m MsgDepositShares) ValidateBasic() error {
	if len(m.ClassId) == 0 {
		return ErrEmptyClassID
	}

	if len(m.Id) == 0 {
		return ErrEmptyNFTID
	}

	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", m.Sender)
	}

	if !m.Shares.GT(sdkmath.ZeroInt()) {
		return errors.Wrapf(sdkerrors.ErrInvalidCoins, "Invalid share amount (%s)", m.Shares)
	}

	return nil
}

// GetSigners returns the expected signers for MsgSend.
func (m MsgDepositShares) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{signer}
}

// MsgWithdrawShares implements the Msg.ValidateBasic method.
func (m MsgWithdrawShares) ValidateBasic() error {
	if len(m.ClassId) == 0 {
		return ErrEmptyClassID
	}

	if len(m.Id) == 0 {
		return ErrEmptyNFTID
	}

	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", m.Sender)
	}

	if !m.Shares.GT(sdkmath.ZeroInt()) {
		return errors.Wrapf(sdkerrors.ErrInvalidCoins, "Invalid share amount (%s)", m.Shares)
	}

	return nil
}

// GetSigners returns the expected signers for MsgSend.
func (m MsgWithdrawShares) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{signer}
}

// ValidateBasic implements the Msg.ValidateBasic method.
func (m MsgSponsor) ValidateBasic() error {
	if len(m.ClassId) == 0 {
		return ErrEmptyClassID
	}

	if len(m.Id) == 0 {
		return ErrEmptyNFTID
	}

	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", m.Sender)
	}

	if !m.Coin.Amount.GT(sdkmath.ZeroInt()) {
		return errors.Wrapf(sdkerrors.ErrInvalidCoins, "Invalid sponsorship coin (%s)", m.Coin)
	}
	err = sdk.ValidateDenom(m.Coin.Denom)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidCoins, "Invalid denom format (%s)", m.Coin)
	}

	return nil
}

// GetSigners returns the expected signers for MsgSend.
func (m MsgSponsor) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{signer}
}
