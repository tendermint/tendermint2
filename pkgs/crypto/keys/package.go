package keys

import (
	"github.com/tendermint/tendermint2/pkgs/amino"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/tendermint/tendermint2/pkgs/crypto/keys",
	"tm.keys",
	amino.GetCallersDirname(),
).WithDependencies().WithTypes(
	localInfo{}, "LocalInfo",
	ledgerInfo{}, "LedgerInfo",
	offlineInfo{}, "OfflineInfo",
	multiInfo{}, "MultiInfo",
))
