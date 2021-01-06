package utils

import (
	"testing"
)

func TestParsePrivKeyEmpty(t *testing.T) {

	_, err := ParsePrivKey("")

	if err == nil {
		t.Error("Expected error")
	}

}

func TestParsePrivKeyBad(t *testing.T) {

	_, err := ParsePrivKey("testdata/bad.key")

	if err == nil {
		t.Error("Expected error")
	}

}

func TestParsePrivKey(t *testing.T) {

	privKey, err := ParsePrivKey("testdata/auth.key")

	if err != nil {
		t.Error("ParsePrivKey failed")
	}

	err = privKey.Validate()

	if err != nil {
		t.Error("Validate failed")
	}

}

func TestParsePubKeyEmpty(t *testing.T) {

	_, err := ParsePubKey("")

	if err == nil {
		t.Error("Expected error")
	}

}

func TestParsePubKeyBad(t *testing.T) {

	_, err := ParsePubKey("testdata/bad.key")

	if err == nil {
		t.Error("Expected error")
	}

}

func TestParsePubKey(t *testing.T) {

	pubKey, err := ParsePubKey("testdata/auth.pub")

	if err != nil {
		t.Error("ParsePrivKey failed")
	}

	if pubKey.E != 65537 {
		t.Error("E should be 65537")
	}

}
