package multisig

import (
	"github.com/tendermint/tendermint2/pkgs/amino"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/tendermint/tendermint2/pkgs/crypto/multisig",
	"tm",
	amino.GetCallersDirname(),
).WithDependencies().WithTypes(
	PubKeyMultisigThreshold{}, "PubKeyMultisig",
))
