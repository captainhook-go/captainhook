package configuration

import (
	"testing"
)

func TestOptionsStringValue(t *testing.T) {
	opt := Options{options: map[string]interface{}{"foo": "bar", "fiz": true, "baz": 10}}

	expected := "bar"
	result := opt.AsString("foo", "")

	if result != expected {
		t.Errorf("Could not find option, got: %s, want: %s.", result, expected)
	}
}

func TestOptionsBoolValue(t *testing.T) {
	opt := Options{options: map[string]interface{}{"foo": "bar", "fiz": true, "baz": 10}}

	expected := true
	result := opt.AsBool("fiz", false)

	if result != expected {
		t.Errorf("Could not find option, got: %t, want: %t.", result, expected)
	}
}

func TestOptionsDefaultValue(t *testing.T) {
	opt := Options{options: map[string]interface{}{"foo": "bar", "fiz": true, "baz": 10}}

	expected := "default"
	result := opt.AsString("not-there", "default")

	if result != expected {
		t.Errorf("Could not find option, got: %s, want: %s.", result, expected)
	}
}
