package lib

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/OdaDaisuke/pssh-parser/pb"
	proto "github.com/golang/protobuf/proto"
)

const WIDEVINE_SYSTEM_ID = "edef8ba979d64acea3c827dcd51d21ed"

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
	DRMName     string

	// Size of PSSH Data part
	DataSize int64

	// PSSH Data body
	DataHex string
	DataRaw []byte
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
	valid, drmName := p.validateSystemID(drmSystem)
	if !valid {
		log.Printf("Unsupported DRM %s\n", drmSystem)
	}

	dataSize, err := strconv.ParseInt(p.HexBin[56:64], 16, 64)
	if err != nil {
		log.Println("hex int cast error")
	}
	dataHex := p.HexBin[64:]
	dataRaw := p.Data[64:]

	p.Summary = &PSSHSummary{
		Type:        psshType,
		SizeHex:     sizeHex,
		SizeDecimal: sizeDecimal,
		Version:     version,
		Flag:        flag,
		DRMSystemID: drmSystem,
		DRMName:     drmName,
		DataSize:    dataSize,
		DataRaw:     dataRaw,
		DataHex:     dataHex,
	}
}

// Print display PSSH summary
func (p *PSSH) Print() {
	fmt.Println("[PSSH Summary]")
	fmt.Println("Size          ", p.Summary.SizeHex)
	fmt.Println("Size(decimal) ", p.Summary.SizeDecimal)
	fmt.Println("Type          ", p.Summary.Type)
	fmt.Println("Version       ", p.Summary.Version)
	fmt.Println("Flag          ", p.Summary.Flag)
	fmt.Println("DRM           ", p.Summary.DRMSystemID)
	fmt.Println("DRM Name      ", p.Summary.DRMName)
	fmt.Println("DataSize      ", p.Summary.DataSize)
	fmt.Println("Data          ", p.Summary.DataHex)

	// TODO: support PlayReady parse
	switch strings.ToLower(p.Summary.DRMSystemID) {
	case WIDEVINE_SYSTEM_ID:
		// parse
		wv := &pb.WidevineCencHeader{}
		if err := proto.Unmarshal(p.Summary.DataRaw, wv); err != nil {
			log.Print("could not unmarshal Widevine proto")
		}
		fmt.Println("Parsed Data  ", wv)
	}
}

// cf. https://dashif.org/identifiers/content_protection/
func (p *PSSH) validateSystemID(s string) (bool, string) {
	// Remove hyphen
	sArr := strings.Split(s, "-")
	s = strings.Join(sArr, "")

	// DRM System ID for only supports DASH
	type SystemID struct {
		ID   string
		Name string
	}
	systemIDs := [...]SystemID{
		{
			ID:   WIDEVINE_SYSTEM_ID,
			Name: "widevine",
		},
		{
			ID:   "9a04f07998404286ab92e65be0885f95",
			Name: "playready",
		},
		{
			ID:   "F239E769EFA348509C16A903C6932EFB",
			Name: "primetime",
		},
		{
			ID:   "45d481cb-8fe0-49c0-ada9-ab2d2455b2f2",
			Name: "corecrypt",
		},
		{
			ID:   "1077efecc0b24d02ace33c1e52e2fb4b",
			Name: "w3c",
		},
		{
			ID:   "6dd8b3c345f44a68bf3a64168d01a4a6",
			Name: "abv",
		},
	}
	for _, sid := range systemIDs {
		if strings.EqualFold(sid.ID, s) {
			return true, sid.Name
		}
	}
	return false, ""
}

func paddingNumber(n string) string {
	b := "00" + n
	return b[len(b)-2:]
}
