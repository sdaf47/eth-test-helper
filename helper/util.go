package helper

import (
	"os"
	"os/exec"
	"log"
	"syscall"
)

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

