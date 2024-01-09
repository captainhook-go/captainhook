package types

import (
	"testing"
)

func TestSubject(t *testing.T) {
	expected := "Foo bar"
	m := NewCommitMessage("Foo bar\n\nFiz baz.", "#")

	if m.Subject() != expected {
		t.Errorf("Wrong subject, expected '" + expected + "' got '" + m.Subject() + "'")
	}
}
