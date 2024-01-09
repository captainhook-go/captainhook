package conditions

import (
	"github.com/captainhook-go/captainhook/test"
	"testing"
)

func TestRestrictionForExternalCondition(t *testing.T) {
	myIO := test.CreateFakeIO()
	conf := test.CreateFakeConfig()
	repo := test.CreateFakeRepo()

	r := NewExternalCommand(myIO, conf, repo)

	if !r.IsApplicableFor("pre-push") {
		t.Errorf("Should be applicabel for pre-push")
	}

	if !r.IsApplicableFor("pre-commit") {
		t.Errorf("Should be applicabel for pre-commit")
	}
}
