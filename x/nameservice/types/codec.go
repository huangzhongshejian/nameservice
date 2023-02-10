package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec for the module
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSetName{}, "nameservice/SetName", nil)
	cdc.RegisterConcrete(MsgBuyName{}, "nameservice/BuyName", nil)
	cdc.RegisterConcrete(MsgDeleteName{}, "nameservice/DeleteName", nil)
}

// // RegisterCodec registers concrete types on codec
// func RegisterCodec(cdc *codec.Codec) {
// 	// TODO: Register the modules msgs
// }

// // ModuleCdc defines the module codec
// var ModuleCdc *codec.Codec

// func init() {
// 	ModuleCdc = codec.New()
// 	RegisterCodec(ModuleCdc)
// 	codec.RegisterCrypto(ModuleCdc)
// 	ModuleCdc.Seal()
// }
