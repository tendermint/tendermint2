package bitarray

import "github.com/tendermint/tendermint2/pkgs/amino"

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/tendermint/tendermint2/pkgs/bitarray",
	"tm",
	amino.GetCallersDirname(),
).WithDependencies().WithTypes(
	BitArray{}, "BitArray",
))
