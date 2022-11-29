package merkle

import (
	amino "github.com/tendermint/tendermint2/pkgs/amino"
)

var cdc *amino.Codec

func init() {
	cdc = amino.NewCodec()
	cdc.Seal()
}
