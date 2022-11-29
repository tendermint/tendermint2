package testutils

import (
	"github.com/tendermint/tendermint2/pkgs/amino"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/tendermint/tendermint2/pkgs/sdk/testutils",
	"sdk.testutils",
	amino.GetCallersDirname(),
).WithDependencies().WithTypes(
	// ...
	&TestMsg{}, "TestMsg",

	// testmsgs.go
	MsgCounter{},
	MsgNoRoute{},
	MsgCounter2{},
))
