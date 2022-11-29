package submodule

import (
	"github.com/tendermint/tendermint2/pkgs/amino/genproto/example/submodule2"
)

type StructSM struct {
	FieldA int
	FieldB string
	FieldC submodule2.StructSM2
}
