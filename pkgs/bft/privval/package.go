package privval

import (
	"github.com/tendermint/tendermint2/pkgs/amino"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/tendermint/tendermint2/pkgs/bft/privval",
	"tm.remotesigner",
	amino.GetCallersDirname(),
).
	WithDependencies().
	WithTypes(

		// Remote Signer
		&PubKeyRequest{},
		&PubKeyResponse{},
		&SignVoteRequest{},
		&SignedVoteResponse{},
		&SignProposalRequest{},
		&SignedProposalResponse{},
		&PingRequest{},
		&PingResponse{},
	))
