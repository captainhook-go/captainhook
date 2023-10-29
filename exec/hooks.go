package exec

import (
	"errors"
	"github.com/captainhook-go/captainhook/config"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/io"
)

func newDefaultHookRunner(hookName string, appIO io.IO, conf *config.Configuration, repo *git.Repository) *HookRunner {
	h := HookRunner{
		hook:         hookName,
		appIO:        appIO,
		config:       conf,
		repo:         repo,
		beforeHook:   func(appIO io.IO, config *config.Configuration, repo *git.Repository) error { return nil },
		beforeAction: func(appIO io.IO, config *config.Action, repo *git.Repository) error { return nil },
		afterAction:  func(appIO io.IO, config *config.Action, repo *git.Repository) error { return nil },
		afterHook:    func(appIO io.IO, config *config.Configuration, repo *git.Repository) error { return nil },
	}
	return &h
}

func NewCommitMsgRunner(appIO io.IO, conf *config.Configuration, repo *git.Repository) *HookRunner {
	return newDefaultHookRunner(hooks.COMMIT_MSG, appIO, conf, repo)
}

func NewPostCheckoutRunner(appIO io.IO, conf *config.Configuration, repo *git.Repository) *HookRunner {
	return newDefaultHookRunner(hooks.POST_CHECKOUT, appIO, conf, repo)
}

func NewPostCommitRunner(appIO io.IO, conf *config.Configuration, repo *git.Repository) *HookRunner {
	return newDefaultHookRunner(hooks.POST_COMMIT, appIO, conf, repo)
}

func NewPostMergeRunner(appIO io.IO, conf *config.Configuration, repo *git.Repository) *HookRunner {
	return newDefaultHookRunner(hooks.POST_MERGE, appIO, conf, repo)
}

func NewPostRewriteRunner(appIO io.IO, conf *config.Configuration, repo *git.Repository) *HookRunner {
	return newDefaultHookRunner(hooks.POST_REWRITE, appIO, conf, repo)
}

func NewPreCommitRunner(appIO io.IO, conf *config.Configuration, repo *git.Repository) *HookRunner {
	return newDefaultHookRunner(hooks.PRE_COMMIT, appIO, conf, repo)
}

func NewPrepareCommitMsgRunner(appIO io.IO, conf *config.Configuration, repo *git.Repository) *HookRunner {
	h := HookRunner{
		hook:   hooks.PREPARE_COMMIT_MSG,
		appIO:  appIO,
		config: conf,
		repo:   repo,
		beforeHook: func(appIO io.IO, config *config.Configuration, repo *git.Repository) error {
			file := appIO.Argument("file")
			if len(file) < 1 {
				return errors.New("file argument is missing")
			}
			return nil
		},
		beforeAction: func(appIO io.IO, config *config.Action, repo *git.Repository) error {
			// load commit message into repo abstraction
			//h.repo.SetCommitMsg(git.NewCommitMessageFromFile(file, commentChar)
			return nil
		},
		afterAction: func(appIO io.IO, config *config.Action, repo *git.Repository) error {
			//
			//if (GetActionType(config.Action()) == "internal") {
			// write commit message to file  h.repo.GetCommitMsg().GetRawContent()
			//}
			return nil
		},
		afterHook: func(appIO io.IO, config *config.Configuration, repo *git.Repository) error {
			return nil
		},
	}
	return &h
}

func NewPrePushRunner(appIO io.IO, conf *config.Configuration, repo *git.Repository) *HookRunner {
	return newDefaultHookRunner(hooks.PRE_PUSH, appIO, conf, repo)
}
