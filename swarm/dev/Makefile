.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/manunit/sciantcoin/cmd/swarm
	go build -o bin/geth     github.com/manunit/sciantcoin/cmd/geth
	go build -o bin/bootnode github.com/manunit/sciantcoin/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/manunit/sciantcoin/swarm/...
