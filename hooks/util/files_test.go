package util

import (
	"testing"
)

func TestContainsAllStrings(t *testing.T) {
	files := []string{"foo", "bar", "baz"}

	if !ContainsAllStrings(files, []string{"foo", "baz"}) {
		t.Errorf("Files should contain 'foo' and 'baz'")
	}
}

func TestDoesNotContainAllStrings(t *testing.T) {
	files := []string{"foo", "bar", "baz"}

	if ContainsAllStrings(files, []string{"foo", "fiz"}) {
		t.Errorf("Files should not contain 'foo' and 'fiz'")
	}
}

func TestContainsAnyStrings(t *testing.T) {
	files := []string{"foo", "bar", "baz"}

	if !ContainsAnyString(files, []string{"foo", "fiz"}) {
		t.Errorf("Files should contain either 'foo' or 'fiz'")
	}
}

func TestDoesNotContainAnyStrings(t *testing.T) {
	files := []string{"foo", "bar", "baz"}

	if ContainsAnyString(files, []string{"fiz", "faz"}) {
		t.Errorf("Files should not contain either 'fiz' or 'faz'")
	}
}
