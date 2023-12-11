package test

import (
	"testing"
)

func TestRunDummyTest(t *testing.T) {
	io := CreateFakeIO()

	if !io.IsQuiet() {
		t.Errorf("Mock IO should be quiet")
	}

	if !io.IsDebug() {
		t.Errorf("Mock IO should have debug verbosity")
	}
}
