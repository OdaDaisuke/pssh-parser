package main

import (
	"encoding/base64"
	"flag"
	"io/ioutil"
	"log"

	lib "github.com/OdaDaisuke/pssh-parser/internal/apps/pssh"
)

func main() {
	var (
		file    string
		str     string
		encoded bool
	)
	flag.StringVar(&file, "f", "", "file flag")
	flag.StringVar(&str, "s", "", "pssh string flag")
	flag.BoolVar(&encoded, "encoded", false, "encoded flag")
	flag.BoolVar(&encoded, "e", false, "encoded flag")
	flag.Parse()
	runMain(file, str, encoded)
}

func runMain(file, psshStr string, encoded bool) {
	var err error
	if file == "" && psshStr == "" {
		log.Fatal("-f or -s must be set")
	}
	var psshBox []byte
	if file != "" {
		dat, err := ioutil.ReadFile(file)
		checkPanic(err)

		if encoded {
			psshBox, err = base64.StdEncoding.DecodeString(string(dat))
			checkPanic(err)
		}
	} else {
		if encoded {
			psshBox, err = base64.StdEncoding.DecodeString(psshStr)
			checkPanic(err)
		}
	}

	pssh := lib.NewPSSH(psshBox)
	pssh.Parse()
	pssh.Print()
}

func checkPanic(err error) {
	if err != nil {
		panic(err)
	}
}
