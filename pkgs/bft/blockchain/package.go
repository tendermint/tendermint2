package blockchain

import (
	"github.com/tendermint/tendermint2/pkgs/amino"
	btypes "github.com/tendermint/tendermint2/pkgs/bft/types"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/tendermint/tendermint2/pkgs/bft/blockchain",
	"tm",
	amino.GetCallersDirname(),
).WithDependencies(
	btypes.Package,
).WithTypes(
	&bcBlockRequestMessage{}, "BlockRequest",
	&bcBlockResponseMessage{}, "BlockResponse",
	&bcNoBlockResponseMessage{}, "NoBlockResponse",
	&bcStatusRequestMessage{}, "StatusRequest",
	&bcStatusResponseMessage{}, "StatusResponse",
))
