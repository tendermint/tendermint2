package mempool

import (
	"github.com/tendermint/tendermint2/pkgs/amino"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/tendermint/tendermint2/pkgs/bft/mempool",
	"tm",
	amino.GetCallersDirname(),
).WithDependencies().WithTypes(
	&TxMessage{},
))
