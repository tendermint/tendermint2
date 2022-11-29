package submodule

import (
	"github.com/tendermint/tendermint2/pkgs/amino"
	"github.com/tendermint/tendermint2/pkgs/amino/genproto/example/submodule2"
)

var Package = amino.RegisterPackage(
	amino.NewPackage(
		"github.com/tendermint/tendermint2/pkgs/amino/genproto/example/submodule",
		"submodule",
		amino.GetCallersDirname(),
	).WithDependencies(
		submodule2.Package,
	).WithTypes(
		StructSM{},
	),
)
