package info

import (
	"testing"
)

func TestTriggersVirtualHook(t *testing.T) {
	expected := "post-change"
	vHook, _ := VirtualHook("post-merge")

	if vHook != expected {
		t.Errorf("Should trigger virtual hook")
	}
}

func TestNoVirtualHook(t *testing.T) {
	_, ok := VirtualHook("commit-message")

	if ok {
		t.Errorf("Should not trigger a virtual hook")
	}
}
