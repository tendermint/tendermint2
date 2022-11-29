package sdk

import (
	"github.com/tendermint/tendermint2/pkgs/amino"
	abci "github.com/tendermint/tendermint2/pkgs/bft/abci/types"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/tendermint/tendermint2/pkgs/sdk",
	"tm",
	amino.GetCallersDirname(),
).
	WithDependencies(
		abci.Package,
	).
	WithTypes(
		Result{},
	))
