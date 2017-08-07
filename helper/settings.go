package helper

import (
	"fmt"
	"os/exec"
	"syscall"
)

type settings struct {
	GenesisPath string
	ChainName   string
	Port        string
	NetworkId   string
	RPCPort     string
	RPCAddr     string
}

// constructor for private type 'settings'
func SettingsConstructor() settings {
	s := settings{}

	return s
}

// remove private chain
func (s *settings) Clear() {
	execCommand(
		"rm",
		"-r",
		"-f",
		s.ChainName,
	)
}

// create new private chain in directory
func (s *settings) Init() {
	syscall.Mkdir(s.ChainName, 0777)

	execCommand(
		"geth",
		"--datadir",
		s.ChainName,
		"init",
		s.GenesisPath,
	)
}

// start private chain
func (s *settings) Run() {
	// TODO how to run?

	cmd := exec.Command(
		"geth",
		"--port",
		s.Port,
		"--networkid",
		s.NetworkId,
		"--nodiscover",
		"--datadir="+s.ChainName,
		"--maxpeers=0",
		"--rpc",
		"-rpcport",
		s.RPCPort,
		"--rpcaddr",
		s.RPCAddr,
		"--rpccorsdomain",
		"\"http://127.0.0.1:8000\"",
		"--rpcapi",
		"\"eth,net,web3\"",
	)

	fmt.Println(cmd.Args)
}
