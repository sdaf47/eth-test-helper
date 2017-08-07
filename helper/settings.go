package helper

import (
	"syscall"
)

type settings struct {
	GenesisPath string
	ChainName string
	GethParams params
}

// constructor for private type 'settings'
func SettingsConstructor() settings {
	s := settings{}
	s.GethParams = params{}

	return s
}

// remove private chain
func (s *settings) Clear() {
	execCommand("rm", []string{
		"rm",
		"-r",
		"-f",
		s.ChainName,
	})
}

// create new private chain in directory
func (s *settings) Start() {
	syscall.Mkdir(s.ChainName, 0777)

	execCommand("geth", []string{
		"geth",
		"--datadir",
		s.ChainName,
		"init",
		s.GenesisPath,
	})
}
