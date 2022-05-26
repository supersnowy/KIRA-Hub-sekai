package types

import (
	"github.com/KiraCore/sekai/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &MsgUpsertStakingPool{}
)

func NewMsgUpsertStakingPool(sender, validator string, enabled bool) *MsgUpsertStakingPool {
	return &MsgUpsertStakingPool{
		Sender:    sender,
		Validator: validator,
		Enabled:   enabled,
	}
}

func (m *MsgUpsertStakingPool) Route() string {
	return ModuleName
}

func (m *MsgUpsertStakingPool) Type() string {
	return types.MsgTypeUpsertStakingPool
}

func (m *MsgUpsertStakingPool) ValidateBasic() error {
	return nil
}

func (m *MsgUpsertStakingPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgUpsertStakingPool) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{
		sender,
	}
}

var (
	_ sdk.Msg = &MsgDelegate{}
)

func NewMsgDelegate(delegator string, validator string, coins sdk.Coins) *MsgDelegate {
	return &MsgDelegate{
		DelegatorAddress: delegator,
		ValidatorAddress: validator,
		Amounts:          coins,
	}
}

func (m *MsgDelegate) Route() string {
	return ModuleName
}

func (m *MsgDelegate) Type() string {
	return types.MsgTypeDelegate
}

func (m *MsgDelegate) ValidateBasic() error {
	return nil
}

func (m *MsgDelegate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgDelegate) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(m.DelegatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{
		sender,
	}
}

var (
	_ sdk.Msg = &MsgUndelegate{}
)

func NewMsgUndelegate(delegator string, validator string, coins sdk.Coins) *MsgUndelegate {
	return &MsgUndelegate{
		DelegatorAddress: delegator,
		ValidatorAddress: validator,
		Amounts:          coins,
	}
}

func (m *MsgUndelegate) Route() string {
	return ModuleName
}

func (m *MsgUndelegate) Type() string {
	return types.MsgTypeUndelegate
}

func (m *MsgUndelegate) ValidateBasic() error {
	return nil
}

func (m *MsgUndelegate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgUndelegate) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(m.DelegatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{
		sender,
	}
}

var (
	_ sdk.Msg = &MsgClaimRewards{}
)

func NewMsgClaimRewards(sender string) *MsgClaimRewards {
	return &MsgClaimRewards{
		Sender: sender,
	}
}

func (m *MsgClaimRewards) Route() string {
	return ModuleName
}

func (m *MsgClaimRewards) Type() string {
	return types.MsgTypeClaimRewards
}

func (m *MsgClaimRewards) ValidateBasic() error {
	return nil
}

func (m *MsgClaimRewards) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgClaimRewards) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{
		sender,
	}
}

var (
	_ sdk.Msg = &MsgClaimRewards{}
)

func NewMsgClaimUndelegation(sender string, undelegationId uint64) *MsgClaimUndelegation {
	return &MsgClaimUndelegation{
		Sender:         sender,
		UndelegationId: undelegationId,
	}
}

func (m *MsgClaimUndelegation) Route() string {
	return ModuleName
}

func (m *MsgClaimUndelegation) Type() string {
	return types.MsgTypeClaimUndelegation
}

func (m *MsgClaimUndelegation) ValidateBasic() error {
	return nil
}

func (m *MsgClaimUndelegation) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgClaimUndelegation) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{
		sender,
	}
}