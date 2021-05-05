package main_test

import (
	"testing"

	main "github.com/iamtio/goradex"
)

//https://github.com/mwwhited/BinaryDataDecoders/blob/master/src/BinaryDataDecoders.Quarta.RadexOne/Notes.md
func TestChecksum(t *testing.T) {
	cs := main.CalcChecksum(0x7bff, 0x2000, 0x0600, 0x4e01)
	if cs != 0x0fff {
		t.Fatalf("Wrong checksum: % X", cs)
	}

	cs = main.CalcChecksum(0x7aff, 0x2080, 0x0600, 0x4e01)
	if cs != 0x107f {
		t.Fatalf("Wrong checksum: %X", cs)
	}
}

func TestLEWord(t *testing.T) {
	if main.LEWord(0xABAC) != 0xACAB {
		t.Fatalf("Can't swap bytes")
	}
	if main.LEWord(0xFF00) != 0x00FF {
		t.Fatalf("Can't swap bytes")
	}
	if main.LEWord(0xABCD) != 0xCDAB {
		t.Fatalf("Can't swap bytes")
	}
	if main.LEWord(0x0001) != 0x0100 {
		t.Fatalf("Can't swap bytes")
	}
	if main.LEWord(0x1001) != 0x0110 {
		t.Fatalf("Can't swap bytes")
	}
	if main.LEWord(0x0101) != 0x0101 {
		t.Fatalf("Can't swap bytes")
	}
}
