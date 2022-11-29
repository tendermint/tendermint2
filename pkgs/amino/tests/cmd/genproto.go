package main

import (
	"github.com/tendermint/tendermint2/pkgs/amino/genproto"
	"github.com/tendermint/tendermint2/pkgs/amino/tests"
)

func main() {
	pkg := tests.Package
	genproto.WriteProto3Schema(pkg)
	genproto.MakeProtoFolder(pkg, "proto")
	genproto.RunProtoc(pkg, "proto")
	genproto.WriteProtoBindings(pkg)
}
