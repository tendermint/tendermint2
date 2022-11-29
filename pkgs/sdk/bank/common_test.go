package bank

// DONTCOVER

import (
	bft "github.com/tendermint/tendermint2/pkgs/bft/types"
	dbm "github.com/tendermint/tendermint2/pkgs/db"
	"github.com/tendermint/tendermint2/pkgs/log"

	"github.com/tendermint/tendermint2/pkgs/sdk"
	"github.com/tendermint/tendermint2/pkgs/sdk/auth"
	"github.com/tendermint/tendermint2/pkgs/std"
	"github.com/tendermint/tendermint2/pkgs/store"
	"github.com/tendermint/tendermint2/pkgs/store/iavl"
)

type testEnv struct {
	ctx  sdk.Context
	bank BankKeeper
	acck auth.AccountKeeper
}

func setupTestEnv() testEnv {
	db := dbm.NewMemDB()

	authCapKey := store.NewStoreKey("authCapKey")

	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(authCapKey, iavl.StoreConstructor, db)
	ms.LoadLatestVersion()

	ctx := sdk.NewContext(sdk.RunTxModeDeliver, ms, &bft.Header{ChainID: "test-chain-id"}, log.NewNopLogger())
	acck := auth.NewAccountKeeper(
		authCapKey, std.ProtoBaseAccount,
	)

	bank := NewBankKeeper(acck)

	return testEnv{ctx: ctx, bank: bank, acck: acck}
}
