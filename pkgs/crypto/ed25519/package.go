package ed25519

import (
	"github.com/tendermint/tendermint2/pkgs/amino"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/tendermint/tendermint2/pkgs/crypto/ed25519",
	"tm",
	amino.GetCallersDirname(),
).WithDependencies().WithTypes(
	PubKeyEd25519{}, "PubKeyEd25519",
	PrivKeyEd25519{}, "PrivKeyEd25519",
))
