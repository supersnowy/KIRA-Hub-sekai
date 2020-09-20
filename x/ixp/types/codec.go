package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterCodec register all the messages for amino
func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateOrder{}, "ixp/MsgCreateOrder", nil)
	cdc.RegisterConcrete(&MsgCancelOrder{}, "ixp/MsgCancelOrder", nil)
	cdc.RegisterConcrete(&MsgCreateOrderBook{}, "ixp/MsgCreateOrderBook", nil)
	cdc.RegisterConcrete(&MsgUpsertSignerKey{}, "ixp/MsgUpsertSignerKey", nil)
}

// RegisterInterfaces register all messages for sdk.Msg type
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateOrder{},
		&MsgCancelOrder{},
		&MsgCreateOrderBook{},
		&MsgUpsertSignerKey{},
	)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/staking module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/staking and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterCodec(amino)
	amino.Seal()
}