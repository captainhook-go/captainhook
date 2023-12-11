package filestaged

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/test"
	"testing"
)

func TestAllFilesAreStaged(t *testing.T) {
	inOut := test.CreateFakeIO()
	conf := test.CreateFakeConfig()
	repo := test.CreateFakeRepo()
	repo.SetFiles([]string{"foo", "bar", "baz"})

	options := configuration.NewOptions(map[string]interface{}{"files": []interface{}{"foo", "bar"}})
	condition := configuration.NewCondition("CaptainHook::StagedFiles.All", options)

	action := NewAll(inOut, conf, repo)
	if !action.IsTrue(condition) {
		t.Errorf("All files should be staged")
	}
}

func TestNotAllFilesAreStaged(t *testing.T) {
	inOut := test.CreateFakeIO()
	conf := test.CreateFakeConfig()
	repo := test.CreateFakeRepo()
	repo.SetFiles([]string{"foo", "bar", "baz"})

	options := configuration.NewOptions(map[string]interface{}{"files": []interface{}{"foo", "fiz"}})
	condition := configuration.NewCondition("CaptainHook::StagedFiles.All", options)

	action := NewAll(inOut, conf, repo)
	if action.IsTrue(condition) {
		t.Errorf("All files should not be staged")
	}
}

func TestAllStagedFilesFailed(t *testing.T) {
	inOut := test.CreateFakeIO()
	conf := test.CreateFakeConfig()
	repo := test.CreateFakeRepo()
	repo.SetFilesError(true)

	options := configuration.NewOptions(map[string]interface{}{"files": []interface{}{"foo", "fiz"}})
	condition := configuration.NewCondition("CaptainHook::StagedFiles.All", options)

	action := NewAll(inOut, conf, repo)
	if action.IsTrue(condition) {
		t.Errorf("Staged files should have failed")
	}
}

func TestAllApplicableFor(t *testing.T) {
	inOut := test.CreateFakeIO()
	conf := test.CreateFakeConfig()
	repo := test.CreateFakeRepo()

	action := NewAll(inOut, conf, repo)
	if !action.IsApplicableFor("pre-commit") {
		t.Errorf("Condition should be applicable for 'pre-commit'")
	}
	if action.IsApplicableFor("pre-push") {
		t.Errorf("Condition should not be applicable for 'pre-push'")
	}
}
