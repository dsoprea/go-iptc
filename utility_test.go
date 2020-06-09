package iptc

import (
	"testing"
)

func TestDumpBytesToString(t *testing.T) {
	s := DumpBytesToString([]byte{1, 2, 3})
	if s != "01 02 03" {
		t.Fatalf("String not correct: [%s]", s)
	}
}
