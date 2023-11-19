package io

import (
	"strings"
	"testing"
)

func TestPrefixLines(t *testing.T) {
	lines := []string{"foo", "bar", "baz"}

	prefixed := PrefixLines(lines, "#")

	for i, line := range prefixed {
		if !strings.HasPrefix(line, "#") {
			t.Errorf("Line should be prefixed, got: %s expected: %s", line, "#"+lines[i])
		}
	}
}
