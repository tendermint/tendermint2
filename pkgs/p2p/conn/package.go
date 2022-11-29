package conn

import (
	"reflect"

	"github.com/tendermint/tendermint2/pkgs/amino"
	"github.com/tendermint/tendermint2/pkgs/amino/pkg"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/tendermint/tendermint2/pkgs/p2p/conn",
	"p2p", // keep short, do not change.
	amino.GetCallersDirname(),
).
	WithDependencies(
	// NA
	).
	WithTypes(

		// NOTE: Keep the names short.
		pkg.Type{
			Type: reflect.TypeOf(PacketPing{}),
			Name: "Ping",
		},
		pkg.Type{
			Type: reflect.TypeOf(PacketPong{}),
			Name: "Pong",
		},
		pkg.Type{
			Type: reflect.TypeOf(PacketMsg{}),
			Name: "Msg",
		},
	))
