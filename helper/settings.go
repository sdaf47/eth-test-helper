package helper

import (
	"syscall"
	"log"
	"os/exec"
	"os"
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

func execCommand(command string, args []string) {
	env := os.Environ()

	binary, err := exec.LookPath(command)
	if err != nil {
		log.Print(err)
	}

	err = syscall.Exec(binary, args, env)
	if err != nil {
		log.Print(err)
	}
}
