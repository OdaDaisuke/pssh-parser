package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	lib "github.com/OdaDaisuke/pssh-parser/lib"
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

	fmt.Println("psshSize", pssh.Summary.SizeHex)
	fmt.Println("psshZie(decimal)", pssh.Summary.SizeDecimal)
	fmt.Println("type", pssh.Summary.Type)
	fmt.Println("version", pssh.Summary.Version)
	fmt.Println("flag", pssh.Summary.Flag)
	fmt.Println("drm", pssh.Summary.DRMSystemID)
	fmt.Println("psshDataSize", pssh.Summary.DataSize)
	fmt.Println("psshData", pssh.Summary.Data)
}

func checkPanic(err error) {
	if err != nil {
		panic(err)
	}
}
