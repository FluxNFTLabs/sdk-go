package types

import (
	"errors"
	fnfttypes "github.com/FluxNFTLabs/sdk-go/chain/modules/fnft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	ClassCommissionsProposalType = "ClassCommissionsProposal"
	VerifiersProposalType        = "VerifiersProposal"
)

var _ govtypes.Content = &ClassCommissionsProposal{}
var _ govtypes.Content = &VerifiersProposal{}

func init() {
	govtypes.RegisterProposalType((&ClassCommissionsProposal{}).ProposalType())
	govtypes.RegisterProposalType((&VerifiersProposal{}).ProposalType())
}

func (p *ClassCommissionsProposal) GetTitle() string {
	return p.Title
}

func (p *ClassCommissionsProposal) GetDescription() string {
	return p.Description
}

func (p *ClassCommissionsProposal) ProposalRoute() string {
	return RouterKey
}

func (p *ClassCommissionsProposal) ProposalType() string {
	return ClassCommissionsProposalType
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *ClassCommissionsProposal) ValidateBasic() error {
	if len(p.ClassCommissions) == 0 {
		return errors.New("commissions array cannot be empty")
	}
	for _, cc := range p.ClassCommissions {
		if len(cc.ClassId) == 0 {
			return fnfttypes.ErrEmptyClassID
		}
		if cc.CommissionDiv == 0 || cc.CommissionMul == 0 {
			return errors.New("commission parts cannot be 0")
		}
	}
	return govtypes.ValidateAbstract(p)
}

func (p *VerifiersProposal) GetTitle() string {
	return p.Title
}

func (p *VerifiersProposal) GetDescription() string {
	return p.Description
}

func (p *VerifiersProposal) ProposalRoute() string {
	return RouterKey
}

func (p *VerifiersProposal) ProposalType() string {
	return VerifiersProposalType
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *VerifiersProposal) ValidateBasic() error {
	if len(p.Verifiers) == 0 {
		return errors.New("verifier list cannot be empty")
	}
	for _, v := range p.Verifiers {
		_, err := sdk.AccAddressFromBech32(v)
		if err != nil {
			return err
		}
	}
	return govtypes.ValidateAbstract(p)
}
