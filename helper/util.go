package helper

import (
	"bytes"
	"fmt"
	"os/exec"
)

func execCommand(command string, args ...string) (err error) {
	cmd := exec.Command(
		command,
		args...,
	)
	fmt.Println(cmd.Args)
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err = cmd.Run()
	if err != nil {
		return err
	}
	for {
		op, err := outb.ReadString('\n')
		if op == "" {
			break
		}
		if err != nil {
			return err
		}
		fmt.Print(op)
	}
	for {
		op, err := errb.ReadString('\n')
		if op == "" {
			break
		}
		if err != nil {
			return err
		}
		fmt.Print(op)
	}

	return nil
}
