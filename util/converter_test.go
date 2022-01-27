package util

import (
	"testing"
)

func TestUtil_B64FromUint32(t *testing.T) {
	encoded := B64FromUint32(1)

	if encoded != "MQ==" {
		t.Errorf("Encoding incorrect, want: 'MQ==', got: '%v'.", encoded)
	}
}

func TestUtil_Uint32FromB64(t *testing.T) {
	decoded, err := Uint32FromB64("MQ==")
	if err != nil {
		t.Errorf("Error occurred when decoding Base64 string, want: 1, got: %v.", err)
	}
	if decoded != 1 {
		t.Errorf("Encoding incorrect, want: 1, got: %v.", decoded)
	}
}

func TestUtil_Uint32FromB64_InvalidInput(t *testing.T) {
	_, err := Uint32FromB64("$1a")
	if err == nil {
		t.Errorf("Error was not caught when decoding an improper base64 string: %v.", err)
	}

	_, err = Uint32FromB64("YWJj")
	if err == nil {
		t.Errorf("Error was not caught when trying to cast string: %v to uint32.", err)
	}
}
