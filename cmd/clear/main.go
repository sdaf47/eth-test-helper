package main

import (
	"flag"
	"os"
	"github.com/sdaf47/eth-test-helper/helper"
)

func main() {

	s := helper.SettingsConstructor()

	fs := flag.NewFlagSet("Flags ", flag.ExitOnError)
	fs.StringVar(&s.ChainName, "chain-name", "privchain", "Name of private chain")

	fs.Parse(os.Args[1:])

	s.Clear()

}
