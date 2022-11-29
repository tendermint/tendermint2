package cstypes

import (
	"github.com/tendermint/tendermint2/pkgs/amino"
	abci "github.com/tendermint/tendermint2/pkgs/bft/abci/types"
	btypes "github.com/tendermint/tendermint2/pkgs/bft/types"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/tendermint/tendermint2/pkgs/bft/consensus/types",
	"tm",
	amino.GetCallersDirname(),
).
	WithGoPkgName("cstypes").
	WithDependencies(
		abci.Package,
		btypes.Package,
	).
	WithTypes(

		// Round state types
		&RoundState{},
		HRS{},
		RoundStateSimple{},
		PeerRoundState{},

		// Misc
		HeightVoteSet{},

		// Event types
		EventNewRoundStep{},
		EventNewValidBlock{},
		EventNewRound{},
		EventCompleteProposal{},
		EventTimeoutPropose{},
		EventTimeoutWait{},
	))
