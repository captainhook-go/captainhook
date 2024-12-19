package status

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/test"
	"testing"
)

func TestBranchConfigMissing(t *testing.T) {
	inOut := test.CreateFakeIO()
	conf := test.CreateFakeConfig()
	repo := test.CreateFakeRepo()

	options := configuration.NewOptions(map[string]interface{}{})
	condition := configuration.NewCondition("CaptainHook::Status.OnBranch", options, []*configuration.Condition{})

	action := NewOnBranch(inOut, conf, repo)
	if action.IsTrue(condition) {
		t.Errorf("Condition should not apply on config error")
	}
}

func TestBranchMainAllowed(t *testing.T) {
	inOut := test.CreateFakeIO()
	conf := test.CreateFakeConfig()
	repo := test.CreateFakeRepo()

	options := configuration.NewOptions(map[string]interface{}{"name": "main"})
	condition := configuration.NewCondition("CaptainHook::Status.OnBranch", options, []*configuration.Condition{})

	action := NewOnBranch(inOut, conf, repo)
	if !action.IsTrue(condition) {
		t.Errorf("Condition should apply")
	}
}

func TestBranchMainNotAllowed(t *testing.T) {
	inOut := test.CreateFakeIO()
	conf := test.CreateFakeConfig()
	repo := test.CreateFakeRepo()

	options := configuration.NewOptions(map[string]interface{}{"name": "foo"})
	condition := configuration.NewCondition("CaptainHook::Status.OnBranch", options, []*configuration.Condition{})

	action := NewOnBranch(inOut, conf, repo)
	if action.IsTrue(condition) {
		t.Errorf("Condition should not apply on wrong branch")
	}
}

func TestApplicableFor(t *testing.T) {
	inOut := test.CreateFakeIO()
	conf := test.CreateFakeConfig()
	repo := test.CreateFakeRepo()

	action := NewOnBranch(inOut, conf, repo)
	if !action.IsApplicableFor("pre-commit") {
		t.Errorf("Condition should be applicable for 'pre-commit'")
	}
	if !action.IsApplicableFor("pre-push") {
		t.Errorf("Condition should be applicable for 'pre-push'")
	}
}
