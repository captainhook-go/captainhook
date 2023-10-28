package exec

import (
	"errors"
	"github.com/captainhook-go/captainhook/config"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/io"
)

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
