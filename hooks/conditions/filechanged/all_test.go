package filechanged

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/test"
	"testing"
)

func TestAllFilesAreChanged(t *testing.T) {
	inOut := test.CreateFakeIO()
	inOut.SetArguments(map[string]string{"command": "pre-push"})
	inOut.SetOptions(map[string]string{"input": "refs/heads/main 12345 refs/heads/main 09876"})
	conf := test.CreateFakeConfig()
	repo := test.CreateFakeRepo()
	repo.SetFiles([]string{"foo", "bar", "baz"})

	options := configuration.NewOptions(map[string]interface{}{"files": []interface{}{"foo", "bar"}})
	condition := configuration.NewCondition("CaptainHook::ChangedFiles.All", options, []*configuration.Condition{})

	action := NewAll(inOut, conf, repo)
	if !action.IsTrue(condition) {
		t.Errorf("All files should be changed")
	}
}

func TestNotAllChangedFiles(t *testing.T) {
	inOut := test.CreateFakeIO()
	inOut.SetArguments(map[string]string{"command": "pre-push"})
	inOut.SetOptions(map[string]string{"input": "refs/heads/main 12345 refs/heads/main 09876"})
	conf := test.CreateFakeConfig()
	repo := test.CreateFakeRepo()
	repo.SetFiles([]string{"foo", "bar", "baz"})

	options := configuration.NewOptions(map[string]interface{}{"files": []interface{}{"foo", "fiz"}})
	condition := configuration.NewCondition("CaptainHook::ChangedFiles.All", options, []*configuration.Condition{})

	action := NewAll(inOut, conf, repo)
	if action.IsTrue(condition) {
		t.Errorf("All files should not be changed")
	}
}

func TestAllChangedFilesDetectionFailed(t *testing.T) {
	inOut := test.CreateFakeIO()
	inOut.SetArguments(map[string]string{"command": "pre-push"})
	inOut.SetOptions(map[string]string{"input": ""})
	conf := test.CreateFakeConfig()
	repo := test.CreateFakeRepo()
	repo.SetFilesError(true)

	options := configuration.NewOptions(map[string]interface{}{"files": []interface{}{"foo", "fiz"}})
	condition := configuration.NewCondition("CaptainHook::ChangedFiles.All", options, []*configuration.Condition{})

	action := NewAll(inOut, conf, repo)
	if action.IsTrue(condition) {
		t.Errorf("Range detection should have failed")
	}
}

func TestAllChangedFilesFailed(t *testing.T) {
	inOut := test.CreateFakeIO()
	inOut.SetArguments(map[string]string{"command": "pre-push"})
	inOut.SetOptions(map[string]string{"input": "refs/heads/main 12345 refs/heads/main 09876"})
	conf := test.CreateFakeConfig()
	repo := test.CreateFakeRepo()
	repo.SetFilesError(true)

	options := configuration.NewOptions(map[string]interface{}{"files": []interface{}{"foo", "fiz"}})
	condition := configuration.NewCondition("CaptainHook::ChangedFiles.All", options, []*configuration.Condition{})

	action := NewAll(inOut, conf, repo)
	if action.IsTrue(condition) {
		t.Errorf("All files should not be changed")
	}
}

func TestAllApplicableFor(t *testing.T) {
	inOut := test.CreateFakeIO()
	conf := test.CreateFakeConfig()
	repo := test.CreateFakeRepo()

	action := NewAll(inOut, conf, repo)
	if action.IsApplicableFor("pre-commit") {
		t.Errorf("Condition should noz be applicable for 'pre-commit'")
	}
	if !action.IsApplicableFor("pre-push") {
		t.Errorf("Condition should be applicable for 'pre-push'")
	}
}
