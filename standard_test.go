package iptc

import (
	"testing"

	"github.com/dsoprea/go-logging"
)

func TestGetTagInfo_Hit(t *testing.T) {
	sti, err := GetTagInfo(9, 10)
	log.PanicIf(err)

	if sti.Description != "Confirmed ObjectData Size" {
		t.Fatalf("Tag description not correct: %s", sti.Description)
	}
}

func TestGetTagInfo_Miss(t *testing.T) {
	_, err := GetTagInfo(99, 99)
	if err != nil && err != ErrTagNotStandard {
		log.PanicIf(err)
	} else if err == nil {
		t.Fatalf("Expected error for invalid tag.")
	}
}
