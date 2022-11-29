package secp256k1

import (
	"github.com/tendermint/tendermint2/pkgs/amino"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/tendermint/tendermint2/pkgs/crypto/secp256k1",
	"tm",
	amino.GetCallersDirname(),
).WithDependencies().WithTypes(
	PubKeySecp256k1{}, "PubKeySecp256k1",
	PrivKeySecp256k1{}, "PrivKeySecp256k1",
))
