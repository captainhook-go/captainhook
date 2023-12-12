package util

import (
	"testing"
)

func TestFilterByType(t *testing.T) {
	files := []string{"foo.txt", "bar.jpg", "baz.gif"}

	filtered := FilterByType(files, ".gif")
	if len(filtered) != 1 {
		t.Errorf("Filter should remove all but one")
	}
}

func TestFilterByTypeWithDir(t *testing.T) {
	files := []string{"foo/foo.txt", "bar/bar.jpg", "baz/baz.gif"}

	filtered := FilterByType(files, ".gif")
	if len(filtered) != 1 {
		t.Errorf("Filter should remove all but one")
	}
}

func TestFilterByDir(t *testing.T) {
	files := []string{"foo/fiz/foo.txt", "foo/bar/bar.jpg", "foo/baz/baz.gif"}

	filtered := FilterByDirectory(files, "foo/fiz")
	if len(filtered) != 1 {
		t.Errorf("Filter should remove all but one")
	}
}

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
