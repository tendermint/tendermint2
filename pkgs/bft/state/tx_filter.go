package state

import (
	mempl "github.com/tendermint/tendermint2/pkgs/bft/mempool"
	"github.com/tendermint/tendermint2/pkgs/bft/types"
)

// TxPreCheck returns a function to filter transactions before processing.
// The function does nothing yet.
func TxPreCheck(state State) mempl.PreCheckFunc {
	return func(types.Tx) error { return nil }
}
