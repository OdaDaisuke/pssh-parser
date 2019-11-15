package lib

import "testing"

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
