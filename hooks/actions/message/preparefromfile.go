package message

import (
	"errors"
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/events"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/git/types"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
	"os"
)

// PrepareFromFile loads a suggested commit message from a given file.
// It is the counterpart of CacheOnFail. If your commit failed because of an
// invalid message you can use the action to preload the failed message.
// This way you can edit it without having to completely write it from scratch again.
// If the current commit message is not empty it will add it the cached message as
// comment at the bottom.
//
// Example configuration:
//
//	{
//	  "run": "CaptainHook::Message.PrepareFromFile",
//	  "options": {
//	    "file": ".git/CH_COMMIT_MSG"
//	  }
//	}
type PrepareFromFile struct {
	hookBundle *hooks.HookBundle
}

func (a *PrepareFromFile) IsApplicableFor(hook string) bool {
	return a.hookBundle.Restriction.IsApplicableFor(hook)
}

func (a *PrepareFromFile) Run(action *configuration.Action) error {
	a.hookBundle.AppIO.Write("prepare from file", true, io.DEBUG)
	gitMsgFile := a.hookBundle.AppIO.Argument(info.ArgCommitMsgFile, "")
	msg, loadErr := a.hookBundle.Repo.CommitMessage(gitMsgFile)
	if loadErr != nil {
		return loadErr
	}

	preparedMessageFile := action.Options().AsString("file", "")
	if preparedMessageFile == "" {
		return errors.New("option 'file' is missing")
	}

	content, readErr := io.ReadFile(preparedMessageFile)
	if readErr != nil {
		a.hookBundle.AppIO.Write("no file to read", true, io.DEBUG)
		return nil
	}
	return a.hookBundle.Repo.PrepareCommitMessage(gitMsgFile, a.createCommitMessage(msg, string(content)))
}

func (a *PrepareFromFile) Subscribe(dispatcher *events.Dispatcher, action *configuration.Action) error {
	a.hookBundle.AppIO.Write("register hook success event", true, io.VERBOSE)
	path := action.Options().AsString("file", "")
	if path == "" {
		return errors.New("option 'file' is missing")
	}
	dispatcher.RegisterHookSucceededSubscribers(NewPrepareFromFileEventHandler(a.hookBundle, path))
	return nil
}

func (a *PrepareFromFile) createCommitMessage(msg *types.CommitMessage, content string) *types.CommitMessage {
	if msg.IsEmpty() {
		// there could be comments so append them to the prepared message
		return types.NewCommitMessage(content+"\n"+msg.Raw(), msg.CommentChar())
	}
	// append the loaded content stuff as a comment
	return types.NewCommitMessage(
		msg.Raw()+"\n"+io.PrefixLinesInString(content, msg.CommentChar()+" "),
		msg.CommentChar(),
	)
}

func NewPrepareFromFile(appIO io.IO, conf *configuration.Configuration, repo git.Repo) hooks.Action {
	a := PrepareFromFile{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.PrepareCommitMsg}),
	}
	return &a
}

type PrepareFromFileEventHandler struct {
	bundle *hooks.HookBundle
	path   string
}

func (h *PrepareFromFileEventHandler) Handle(event *events.HookSucceeded) error {
	h.bundle.AppIO.Write("  PrepareFromFile - handle succeeded event", true, io.VERBOSE)

	if removeErr := os.Remove(h.path); removeErr != nil {
		h.bundle.AppIO.Write("  PrepareFromFile - could not delete file", true, io.NORMAL)
		return errors.New("could not delete 'PrepareFromFile' file")
	}
	return nil
}

func NewPrepareFromFileEventHandler(bundle *hooks.HookBundle, path string) *PrepareFromFileEventHandler {
	return &PrepareFromFileEventHandler{bundle: bundle, path: path}
}
