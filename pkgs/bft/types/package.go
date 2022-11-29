package types

import (
	"github.com/tendermint/tendermint2/pkgs/amino"
	abci "github.com/tendermint/tendermint2/pkgs/bft/abci/types"
	"github.com/tendermint/tendermint2/pkgs/bitarray"
	"github.com/tendermint/tendermint2/pkgs/crypto/merkle"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/tendermint/tendermint2/pkgs/bft/types",
	"tm",
	amino.GetCallersDirname(),
).
	WithDependencies(
		abci.Package,
		bitarray.Package,
		merkle.Package,
	).
	WithTypes(

		// Proposal
		Proposal{},

		// Block types
		Block{},
		Header{},
		Data{},
		// EvidenceData{},
		Commit{},
		BlockID{},
		CommitSig{},
		Vote{},
		// Tx{},
		// Txs{},
		Part{},
		PartSet{},
		PartSetHeader{},

		// Internal state types
		Validator{},
		ValidatorSet{},

		// Event types
		EventNewBlock{},
		EventNewBlockHeader{},
		EventTx{},
		EventVote{},
		EventString(""),
		EventValidatorSetUpdates{},

		// Evidence types
		DuplicateVoteEvidence{},
		MockGoodEvidence{},
		MockRandomGoodEvidence{},
		MockBadEvidence{},

		// Misc.
		TxResult{},
		MockAppState{},
		VoteSet{},
	))
