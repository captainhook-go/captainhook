package exec

import (
	"testing"
)

func TestDetectInternal(t *testing.T) {
	if !isInternalFunctionality("CaptainHook::") {
		t.Errorf("Should be internal, tested: %s", "CaptainHook::")
	}
	if isInternalFunctionality("my/captainhook/stuff") {
		t.Errorf("Should not be internal, tested: %s", "my/captainhook/stuff")
	}
	if isInternalFunctionality("something::with::colons") {
		t.Errorf("Should not be internal, tested: %s", "something::with::colons")
	}
}
