package main

import (
	"flag"
	"github.com/sdaf47/eth-test-helper/helper"
	"os"
)

const (
	ActionStart = "start"
	ActionClear = "clear"
)

func main() {

	s := helper.SettingsConstructor()
	var action string

	fs := flag.NewFlagSet("Flags ", flag.ExitOnError)
	fs.StringVar(&s.GenesisPath, "genesis-path", "genesis.json", "Path to genesis")
	fs.StringVar(&s.ChainName, "chain-name", "privchain", "Name of private chain")
	fs.StringVar(&s.Port, "port", "3000", "Port for geth listen")
	fs.StringVar(&s.NetworkId, "networkid", "58342", "Id of network")
	fs.StringVar(&s.RPCPort, "rpcport", "8545", "Port for Ethereum rpc access")
	fs.StringVar(&s.RPCAddr, "rpcaddr", "127.0.0.1", "Host for rpc running")
	fs.StringVar(&action, "action", ActionStart, "Action")

	fs.Parse(os.Args[1:])

	switch action {
	case ActionStart:
		s.Clear()
		s.Init()
		s.Run()
	case ActionClear:
		s.Clear()
	}

}
