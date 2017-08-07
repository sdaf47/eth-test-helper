package main

import (
	"flag"
	"github.com/sdaf47/eth-test-helper/helper"
	"os"
)

func main() {
	var loaderPath string

	fs := flag.NewFlagSet("Flags ", flag.ExitOnError)
	fs.StringVar(&loaderPath, "file", "Contract.sol", "Contract file path")

	fs.Parse(os.Args[1:])

	l := helper.LoaderConstructor(loaderPath)
	l.Make()

}
