package frames

import (
	"testing"
)

func TestFirst(t *testing.T) {
	var ti Info
	ti.Init("/dev/null")
}

func TestDecode(t *testing.T) {
	expected := map[string]string{
		"MOTDETAT": "000000",
		"ADCO":     "020828337598",
		"OPTARIF":  "BBR(",
		"ISOUSC":   "45",
	}
	got := make(map[string]string)
	gotframe := make(map[string]string)
	cb := func(code string, value string) {
		got[code] = value
	}
	framecb := func(data map[string]string) {
		gotframe = data
	}
	var ti Info
	ti.Init("/dev/null")
	ti.SetFieldCB(cb)
	ti.SetFrameCB(framecb)

	frame := "\002MOTDETAT 000000 B\nADCO 020828337598 N\nOPTARIF BBR( S\nISOUSC 45 ?\nBBRHCJB 012133887 >\nBBRHPJB 038554302 H\nBBRHCJW 002903317 K\nBBRHPJW 003800290 U\nBBRHCJR 001504374 E\nBBRHPJR 000907447 Y\nPTEC HPJB P\nDEMAIN BLEU V\nIINST 004 [\nIMAX 049 L\nPAPP 01100 #\nHHPHC Y D\n\003"
	for i := 0; i < len(frame); i++ {
		ti.decode(frame[i])
	}
	for code, value := range expected {
		if got[code] != value {
			t.Error("Value for ", code, " is ", got[code], " instead of ", value)
		}
		if gotframe[code] != value {
			t.Error("Frame value for ", code, " is ", gotframe[code], " instead of ", value)
		}
	}
}
