package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/sdaf47/eth-test-helper/helper"
)

func main() {

	s := &helper.SettingsConstructor()

	fs := flag.NewFlagSet("Flags ", flag.ExitOnError)
	fs.StringVar(&s.GenesisPath, "genesis-path", "genesis.json", "Path to genesis")
	fs.StringVar(&s.ChainName, "chain-name", "privchain", "Name of private chain")
	fs.IntVar(&s.GethParams.Port, "port", 3000, "Port for geth listen")
	fs.IntVar(&s.GethParams.NetworkId, "networkid", 58342, "Id of network")
	fs.IntVar(&s.GethParams.RPCPort, "rpcport", 8545, "Port for Ethereum rpc access")
	fs.StringVar(&s.GethParams.RPCAddr, "rpcaddr", "127.0.0.1", "Host for rpc running")

	fs.Parse(os.Args[1:])

	fmt.Println(s)
}
