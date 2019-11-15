package lib

import (
	"encoding/base64"
	"testing"
)

func TestPaddingNumber(t *testing.T) {
	s := "5"
	if paddingNumber(s) != "05" {
		t.Error()
	}

	s = "005"
	if paddingNumber(s) != "05" {
		t.Error()
	}
}

func TestParse(t *testing.T) {
	testDataBin := "AAAAxnBzc2gBAAAA7e+LqXnWSs6jyCfc1R0h7QAAAAINw+xPdoNUi4HnPGTlguE2FEe37S9mVyu9EwbOfPNhDQAAAIISEBRHt+0vZlcrvRMGznzzYQ0SEFrGoR6qL17Vv2aMQByBNMoSEG7hNRbI51h7rp9+zT6Zom4SEPnsEqYaJl1Hj4MzTjp40scSEA3D7E92g1SLgec8ZOWC4TYaDXdpZGV2aW5lX3Rlc3QiEXVuaWZpZWQtc3RyZWFtaW5nSOPclZsG"
	psshBox, _ := base64.StdEncoding.DecodeString(testDataBin)
	pssh := NewPSSH(psshBox)
	pssh.Parse()
	pssh.Print()
	if pssh.Summary == nil {
		t.Error("pssh summary must not be nil")
	}

	if pssh.Summary.SizeHex != "000000c6" {
		t.Errorf("Summary.SizeHex must be 000000c6 got %s", pssh.Summary.SizeHex)
	}
	if pssh.Summary.SizeDecimal != 198 {
		t.Errorf("Summary.SizeHex must be 198 got %d", pssh.Summary.SizeDecimal)
	}
	if pssh.Summary.Type != "70737368" {
		t.Errorf("Summary.Type must be 70737368 got %s", pssh.Summary.Type)
	}
	if pssh.Summary.Version != "01" {
		t.Errorf("Summary.Version must be 01 got %s", pssh.Summary.Version)
	}
	if pssh.Summary.Flag != "000000" {
		t.Errorf("Summary.Flag must be 000000 got %s", pssh.Summary.Flag)
	}
	if pssh.Summary.DRMSystemID != "edef8ba979d64acea3c827dcd51d21ed(widevine)" {
		t.Errorf("Summary.DRMSystemID must be edef8ba979d64acea3c827dcd51d21ed(widevine) got %s", pssh.Summary.DRMSystemID)
	}
	if pssh.Summary.DataSize != 2 {
		t.Errorf("Summary.Data must be 70737368 got %s", pssh.Summary.Type)
	}

	dataExpected := "0dc3ec4f7683548b81e73c64e582e1361447b7ed2f66572bbd1306ce7cf3610d0000008212101447b7ed2f66572bbd1306ce7cf3610d12105ac6a11eaa2f5ed5bf668c401c8134ca12106ee13516c8e7587bae9f7ecd3e99a26e1210f9ec12a61a265d478f83334e3a78d2c712100dc3ec4f7683548b81e73c64e582e1361a0d7769646576696e655f746573742211756e69666965642d73747265616d696e6748e3dc959b06"
	if pssh.Summary.Data != dataExpected {
		t.Errorf("Summary.Data must be %s got %s", dataExpected, pssh.Summary.Type)
	}
}
