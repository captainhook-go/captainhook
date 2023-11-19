package commands

import (
	"testing"
)

func TestOptionsStringValue(t *testing.T) {

	args := []string{"foo", "bar"}
	mapping := []string{"FIRST", "SECOND"}

	mapped := mapArgs(mapping, args, "EXEC")

	first, _ := mapped["FIRST"]
	second, _ := mapped["SECOND"]
	cmd, _ := mapped["command"]

	if first != "foo" {
		t.Errorf("Should be mapped to FIRST, got: %s, want: %s.", first, "foo")
	}
	if second != "bar" {
		t.Errorf("Should be mapped to SECOND, got: %s, want: %s.", second, "bar")
	}
	if cmd != "EXEC" {
		t.Errorf("Should be mapped to FIRST, got: %s, want: %s.", cmd, "exec")
	}
}
