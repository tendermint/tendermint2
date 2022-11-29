package hd

import (
	"github.com/tendermint/tendermint2/pkgs/amino"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/tendermint/tendermint2/pkgs/crypto/hd",
	"tm",
	amino.GetCallersDirname(),
).WithDependencies().WithTypes(
	BIP44Params{}, "Bip44Params",
))
