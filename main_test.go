package main_test

import (
	"testing"

	main "github.com/iamtio/goradex"
)

//https://github.com/mwwhited/BinaryDataDecoders/blob/master/src/BinaryDataDecoders.Quarta.RadexOne/Notes.md
func TestChecksum1(t *testing.T) {
	req := main.MakeQuantaRadexOneRequest(
		0x7bff, 0x2000, 0x0600, 0x4e01,
	)
	req.CalcChecksum0()
	if req.CheckSum0 != 0x0fff {
		t.Fatalf("Wrong checksum: %X", req.CheckSum0)
	}
}
func TestChecksum2(t *testing.T) {
	req := main.MakeQuantaRadexOneRequest(
		0x7aff, 0x2080, 0x0600, 0x4e01,
	)
	req.CalcChecksum0()
	if req.CheckSum0 != 0x107f {
		t.Fatalf("Wrong checksum: %X", req.CheckSum0)
	}
}
