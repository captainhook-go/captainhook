package filechanged

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/test"
	"testing"
)

func TestAnyFilesAreStaged(t *testing.T) {
	inOut := test.CreateFakeIO()
	inOut.SetArguments(map[string]string{"command": "pre-push"})
	inOut.SetOptions(map[string]string{"input": "refs/heads/main 12345 refs/heads/main 09876"})
	conf := test.CreateFakeConfig()
	repo := test.CreateFakeRepo()
	repo.SetFiles([]string{"foo", "bar", "baz"})

	options := configuration.NewOptions(map[string]interface{}{"files": []interface{}{"foo", "fiz"}})
	condition := configuration.NewCondition("CaptainHook::StagedFiles.Any", options, []*configuration.Condition{})

	action := NewAny(inOut, conf, repo)
	if !action.IsTrue(condition) {
		t.Errorf("At least one file should be staged")
	}
}

func TestAnyNoFilesAreStaged(t *testing.T) {
	inOut := test.CreateFakeIO()
	inOut.SetArguments(map[string]string{"command": "pre-push"})
	inOut.SetOptions(map[string]string{"input": "refs/heads/main 12345 refs/heads/main 09876"})
	conf := test.CreateFakeConfig()
	repo := test.CreateFakeRepo()
	repo.SetFiles([]string{"foo", "bar", "baz"})

	options := configuration.NewOptions(map[string]interface{}{"files": []interface{}{"bing", "bang"}})
	condition := configuration.NewCondition("CaptainHook::StagedFiles.Any", options, []*configuration.Condition{})

	action := NewAny(inOut, conf, repo)
	if action.IsTrue(condition) {
		t.Errorf("No files should be staged")
	}
}

func TestAnyStagedFilesFailed(t *testing.T) {
	inOut := test.CreateFakeIO()
	inOut.SetArguments(map[string]string{"command": "pre-push"})
	inOut.SetOptions(map[string]string{"input": "refs/heads/main 12345 refs/heads/main 09876"})
	conf := test.CreateFakeConfig()
	repo := test.CreateFakeRepo()
	repo.SetFilesError(true)

	options := configuration.NewOptions(map[string]interface{}{"files": []interface{}{"foo", "fiz"}})
	condition := configuration.NewCondition("CaptainHook::StagedFiles.Any", options, []*configuration.Condition{})

	action := NewAny(inOut, conf, repo)
	if action.IsTrue(condition) {
		t.Errorf("Staged files should have failed")
	}
}

func TestAnyStagedFilesDetectFailed(t *testing.T) {
	inOut := test.CreateFakeIO()
	inOut.SetArguments(map[string]string{"command": "pre-push"})
	inOut.SetOptions(map[string]string{"input": ""})
	conf := test.CreateFakeConfig()
	repo := test.CreateFakeRepo()

	options := configuration.NewOptions(map[string]interface{}{"files": []interface{}{"foo", "fiz"}})
	condition := configuration.NewCondition("CaptainHook::StagedFiles.Any", options, []*configuration.Condition{})

	action := NewAny(inOut, conf, repo)
	if action.IsTrue(condition) {
		t.Errorf("Range detection should have failed")
	}
}

func TestAnyApplicableFor(t *testing.T) {
	inOut := test.CreateFakeIO()
	conf := test.CreateFakeConfig()
	repo := test.CreateFakeRepo()

	action := NewAny(inOut, conf, repo)
	if action.IsApplicableFor("pre-commit") {
		t.Errorf("Condition should not be applicable for 'pre-commit'")
	}
	if !action.IsApplicableFor("pre-push") {
		t.Errorf("Condition should be applicable for 'pre-push'")
	}
}
