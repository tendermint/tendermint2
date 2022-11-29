package submodule2

import (
	"github.com/tendermint/tendermint2/pkgs/amino"
)

var Package = amino.RegisterPackage(
	amino.NewPackage(
		"github.com/tendermint/tendermint2/pkgs/amino/genproto/example/submodule2",
		"submodule2",
		amino.GetCallersDirname(),
	).WithTypes(
		StructSM2{},
	),
)
