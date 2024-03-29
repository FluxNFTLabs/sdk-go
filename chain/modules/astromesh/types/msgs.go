package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgAstroTransfer{}

func (m *MsgAstroTransfer) ValidateBasic() error {
	if m.Sender == "" {
		return fmt.Errorf("empty sender")
	}

	if m.Receiver == "" {
		return fmt.Errorf("empty receiver")
	}

	if _, exist := Plane_name[int32(m.SrcPlane)]; !exist {
		return fmt.Errorf("unsupported source plane: %d", m.SrcPlane)
	}

	if _, exist := Plane_name[int32(m.DstPlane)]; !exist {
		return fmt.Errorf("unsupported dest plane: %d", m.DstPlane)
	}

	if m.SrcPlane == m.DstPlane {
		return fmt.Errorf("source and dest plane must be different")
	}

	if err := m.Coin.Validate(); err != nil {
		return fmt.Errorf("coin format err: %w", err)
	}

	return nil
}

func (m MsgAstroTransfer) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{signer}
}
