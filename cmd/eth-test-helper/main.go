package main

import (
	"flag"
	"github.com/sdaf47/eth-test-helper/helper"
	"os"
)

const (
	ActionInit  = "init"
	ActionClean = "clean"
)

func main() {

	s := helper.SettingsConstructor()
	var action string
	var loaderPath string

	fs := flag.NewFlagSet("Flags ", flag.ExitOnError)
	fs.StringVar(&s.GenesisPath, "genesis-path", "genesis.json", "Path to genesis")
	fs.StringVar(&s.ChainName, "chain-name", "privchain", "Name of private chain")
	fs.StringVar(&s.Port, "port", "3000", "Port for geth listen")
	fs.StringVar(&s.NetworkId, "networkid", "58342", "Id of network")
	fs.StringVar(&s.RPCPort, "rpcport", "8545", "Port for Ethereum rpc access")
	fs.StringVar(&s.RPCAddr, "rpcaddr", "127.0.0.1", "Host for rpc running")
	fs.StringVar(&loaderPath, "file", "Contract.sol", "Contract file path")
	fs.StringVar(&action, "action", ActionInit, "Action")

	fs.Parse(os.Args[1:])

	switch action {
	case ActionInit:
		s.Clear()
		s.Init()
	case ActionClean:
		s.Clear()
	}

}
