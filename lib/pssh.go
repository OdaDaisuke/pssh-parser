package lib

import (
	"fmt"
	"log"
	"strconv"
)

// PSSH represents mp4 FullBox
type PSSH struct {
	Data    []byte
	HexBin  string
	Summary *PSSHSummary
}

type PSSHSummary struct {
	// PSSH Size attribute
	SizeHex     string
	SizeDecimal int64

	// PSSH Type attribute
	Type string

	// Version attribute
	Version string

	// Flag attbirute
	Flag string

	// DRM System id defined in DASH IF
	DRMSystemID string

	// Size of PSSH Data part
	DataSize int64

	// PSSH Data body
	Data string
}

// NewPSSH generate PSSH
func NewPSSH(p []byte) *PSSH {
	if len(p) < 64 {
		log.Fatal("Invalid data.")
	}
	var hexPssh string
	for _, s := range p {
		sHex := fmt.Sprintf("%x", s)
		hexPssh += paddingNumber(string(sHex))
	}

	return &PSSH{
		Data:   p,
		HexBin: hexPssh,
	}
}

// Parse generate PSSH summary
func (p *PSSH) Parse() {
	psshType := p.HexBin[8:16]
	sizeHex := p.HexBin[0:8]
	sizeDecimal, err := strconv.ParseInt(sizeHex, 16, 64)
	if err != nil {
		log.Println("hex int cast error")
	}
	version := p.HexBin[16:18]
	flag := p.HexBin[18:24]
	drmSystem := p.HexBin[24:56]
	dataSize, err := strconv.ParseInt(p.HexBin[56:64], 16, 64)
	if err != nil {
		log.Println("hex int cast error")
	}
	data := p.HexBin[64:]

	p.Summary = &PSSHSummary{
		Type:        psshType,
		SizeHex:     sizeHex,
		SizeDecimal: sizeDecimal,
		Version:     version,
		Flag:        flag,
		DRMSystemID: drmSystem,
		DataSize:    dataSize,
		Data:        data,
	}
}

func paddingNumber(n string) string {
	b := "00" + n
	return b[len(b)-2:]
}
