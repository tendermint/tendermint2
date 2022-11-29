package main

import (
	"github.com/tendermint/tendermint2/pkgs/amino"
	"github.com/tendermint/tendermint2/pkgs/amino/genproto"

	// TODO: move these out.
	abci "github.com/tendermint/tendermint2/pkgs/bft/abci/types"
	"github.com/tendermint/tendermint2/pkgs/bft/blockchain"
	"github.com/tendermint/tendermint2/pkgs/bft/consensus"
	ctypes "github.com/tendermint/tendermint2/pkgs/bft/consensus/types"
	"github.com/tendermint/tendermint2/pkgs/bft/mempool"
	btypes "github.com/tendermint/tendermint2/pkgs/bft/types"
	"github.com/tendermint/tendermint2/pkgs/bitarray"
	"github.com/tendermint/tendermint2/pkgs/crypto/ed25519"
	"github.com/tendermint/tendermint2/pkgs/crypto/hd"
	"github.com/tendermint/tendermint2/pkgs/crypto/merkle"
	"github.com/tendermint/tendermint2/pkgs/crypto/multisig"
	"github.com/tendermint/tendermint2/pkgs/sdk"
	"github.com/tendermint/tendermint2/pkgs/sdk/bank"
	"github.com/tendermint/tendermint2/pkgs/sdk/vm"
	"github.com/tendermint/tendermint2/pkgs/std"
)

func main() {
	pkgs := []*amino.Package{
		bitarray.Package,
		merkle.Package,
		abci.Package,
		btypes.Package,
		consensus.Package,
		ctypes.Package,
		mempool.Package,
		ed25519.Package,
		blockchain.Package,
		hd.Package,
		multisig.Package,
		std.Package,
		sdk.Package,
		bank.Package,
		vm.Package,
	}
	for _, pkg := range pkgs {
		genproto.WriteProto3Schema(pkg)
		genproto.WriteProtoBindings(pkg)
		genproto.MakeProtoFolder(pkg, "proto")
	}
	for _, pkg := range pkgs {
		genproto.RunProtoc(pkg, "proto")
	}
}
