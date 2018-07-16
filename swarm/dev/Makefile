.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/manunit/go-ethereum-node-sciantcoincmd/swarm
	go build -o bin/geth     github.com/manunit/go-ethereum-node-sciantcoincmd/geth
	go build -o bin/bootnode github.com/manunit/go-ethereum-node-sciantcoincmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/manunit/go-ethereum-node-sciantcoinswarm/...
