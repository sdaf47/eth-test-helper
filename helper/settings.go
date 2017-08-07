package helper

import (
	"syscall"
)

type settings struct {
	GenesisPath string
	ChainName string
	GethParams params
}

func SettingsConstructor() settings {
	s := settings{}
	s.GethParams = params{}

	return s
}

func (s *settings) Clear() {
	execCommand("rm", []string{
		"rm",
		"-r",
		"-f",
		s.ChainName,
	})
}

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
