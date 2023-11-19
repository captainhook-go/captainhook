package git

import (
	"testing"
)

func TestDetectZeroHash(t *testing.T) {
	if IsZeroHash("12345") {
		t.Errorf("Should not be recognized as ZERO-hash, tested: %s", "12345")
	}
	if IsZeroHash("e7ad241a") {
		t.Errorf("Should not be recognized as ZERO-hash, tested: %s", "e7ad241a")
	}
	if !IsZeroHash("0000000000000000000000000000000000000000") {
		t.Errorf("Should be recognized as ZERO-hash, tested: %s", "0000000000000000000000000000000000000000")
	}
}
