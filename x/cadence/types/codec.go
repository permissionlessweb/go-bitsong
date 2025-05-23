package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var (
	amino    = codec.NewLegacyAmino()
	AminoCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)

}

// RegisterLegacyAminoCodec registers concrete types on the LegacyAmino codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegisterCadenceContract{}, "cadence/MsgRegisterCadenceContract", nil)
	cdc.RegisterConcrete(&MsgUnregisterCadenceContract{}, "cadence/MsgUnregisterCadenceContract", nil)
	cdc.RegisterConcrete(&MsgUnjailCadenceContract{}, "cadence/MsgUnjailCadenceContract", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "cadence/MsgUpdateParams", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgRegisterCadenceContract{},
		&MsgUnregisterCadenceContract{},
		&MsgUnjailCadenceContract{},
		&MsgUpdateParams{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
