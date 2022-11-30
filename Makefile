########################################
# Dist suite
.PHONY: reset
all: build

build:
	mkdir -p build
	go build -o ./build/tendermint2 ./cmd/tendermint2

install:
	go install ./cmd/tendermint2

reset:
	rm -rf testdir
	make

clean:
	rm -rf build

########################################
# Formatting, linting.

.PHONY: fmt
fmt:
	go run -modfile ./misc/devdeps/go.mod mvdan.cc/gofumpt -w .

########################################
# Test suite
.PHONY: test test.go test.go1 test.go2 test.go3 test.flappy
test: test.go test.flappy
	@echo "Full test suite finished."

test.flappy:
	# flappy tests should work "sometimes" (at least once)
	TEST_STABILITY=flappy go run -modfile ./misc/devdeps/go.mod moul.io/testman test -test.v -timeout=20m -retry=10 -run ^TestFlappy \
		./pkgs/bft/consensus ./pkgs/bft/blockchain ./pkgs/bft/mempool ./pkgs/p2p ./pkgs/bft/privval

test.go: test.go1 test.go2 test.go3

test.go1:
	# test most of pkgs/* except amino and bft.
	# -p 1 shows test failures as they come
	# maybe another way to do this?
	go test `go list ./pkgs/... | grep -v pkgs/amino/ | grep -v pkgs/bft/` -v -p 1 -timeout=30m

test.go2:
	# test amino.
	go test ./pkgs/amino/... -v -p 1 -timeout=30m

test.go3:
	# test bft.
	go test ./pkgs/bft/... -v -p 1 -timeout=30m

genproto:
	rm -rf proto/*
	-find pkgs | grep -v "^pkgs\/amino" | grep "\.proto" | xargs rm
	-find pkgs | grep -v "^pkgs\/amino" | grep "pbbindings" | xargs rm
	-find pkgs | grep -v "^pkgs\/amino" | grep "pb.go" | xargs rm
	@rm gno.proto || true
	@rm pbbindings.go || true
	@rm gno.pb.go || true
	go run cmd/genproto/genproto.go
	make fmt
