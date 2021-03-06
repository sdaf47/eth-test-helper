package helper

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

const LoaderTemplate = `
var compiled = <COMPILED>;
var binary = "<BINARY>";

var contract = web3.eth.contract(compiled);

var <NAME> = contract.new({from:web3.eth.accounts[0], data: binary, gas: 300000}, function(e, contract){
    if(!e) {

        if(!contract.address) {
            console.log("Contract transaction send: TransactionHash: " + contract.transactionHash + " waiting to be mined...");

        } else {
            console.log("Contract mined! Address: " + contract.address);
            console.log(contract);
        }

    } else {
    	console.log(e);
    }

});
`

type loader struct {
	Path string
	Name string
}

func LoaderConstructor(path string) loader {
	var l loader
	l.Path = path
	// TODO rewrite this shit
	g := regexp.MustCompile("([a-zA-Z0-9]+)[.]sol").FindStringSubmatch(path)
	l.Name = g[1]

	return l
}

func (l *loader) Make() {
	l.makeBin()
	l.makeAbi()

	// TODO and this f**ng shit eto prosto zhest` kto ee pustil

	tmpName := string(regexp.MustCompile("[./]").ReplaceAll([]byte(l.Path), []byte("_")))

	binName := tmpName + "_" + l.Name + ".bin"
	abiName := tmpName + "_" + l.Name + ".abi"
	var binContent []byte
	var abiContent []byte
	var err error

	binContent, err = ioutil.ReadFile(binName)
	if err != nil {
		panic(err)
	}

	abiContent, err = ioutil.ReadFile(abiName)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(l.Name + ".js")
	if err != nil {
		panic(err)
	}

	res := strings.Replace(LoaderTemplate, "<BINARY>", "0x"+string(binContent), 1)
	res = strings.Replace(res, "<COMPILED>", string(abiContent), 1)
	res = strings.Replace(res, "<NAME>", string(l.Name), 1)

	f.Write([]byte(res))
	f.Close()
}

func (l *loader) makeBin() {
	execCommand(
		"solcjs",
		"--bin",
		l.Path,
	)
}

func (l *loader) makeAbi() {
	execCommand(
		"solcjs",
		"--abi",
		l.Path,
	)
}
