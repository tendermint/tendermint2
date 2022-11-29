package bank

import (
	"github.com/tendermint/tendermint2/pkgs/amino"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/tendermint/tendermint2/pkgs/sdk/bank",
	"bank",
	amino.GetCallersDirname(),
).WithDependencies().WithTypes(
	NoInputsError{}, "NoInputsError",
	NoOutputsError{}, "NoOutputsError",
	InputOutputMismatchError{}, "InputOutputMismatchError",
	MsgSend{}, "MsgSend",
))
