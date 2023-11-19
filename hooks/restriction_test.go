package hooks

import (
	"testing"
)

func TestRestrictionForSingleHook(t *testing.T) {
	r := NewRestriction([]string{"pre-commit"})
	if !r.IsApplicableFor("pre-commit") {
		t.Errorf("Should be applicabel for pre-commit")
	}
}

func TestRestrictionForMultipleHooks(t *testing.T) {
	r := NewRestriction([]string{"pre-commit", "pre-push"})
	if !r.IsApplicableFor("pre-commit") {
		t.Errorf("Should be applicabel for pre-commit")
	}
	if !r.IsApplicableFor("pre-push") {
		t.Errorf("Should be applicabel for pre-push")
	}
}

func TestRestrictionNotApplicable(t *testing.T) {
	r := NewRestriction([]string{"pre-commit", "pre-push"})
	if r.IsApplicableFor("post-merge") {
		t.Errorf("Should not be applicabel for post-merge")
	}
	if r.IsApplicableFor("post-commit") {
		t.Errorf("Should not be applicabel for post-commit")
	}
}
