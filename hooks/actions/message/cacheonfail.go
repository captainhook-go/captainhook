package message

import (
	"errors"
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/events"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
	"os"
)

// CacheOnFail writes the commit message to a cache file if the message validation fails.
// This way you can use the cache file to prepare your commit message.
//
// Example configuration:
//
//	{
//	  "run": "CaptainHook::Message.CacheOnFail",
//	  "options: {
//	    "file": ".git/CH_COMMIT_MSG_CACHE"
//	  }
//	}
type CacheOnFail struct {
	hookBundle *hooks.HookBundle
}

func (a *CacheOnFail) IsApplicableFor(hook string) bool {
	return a.hookBundle.Restriction.IsApplicableFor(hook)
}

func (a *CacheOnFail) Run(action *configuration.Action) error {
	a.hookBundle.AppIO.Write(
		"doing nothing just here to register an event:"+action.Options().AsString("file", ""),
		true,
		io.VERBOSE,
	)
	return nil
}

func (a *CacheOnFail) Subscribe(dispatcher *events.Dispatcher, action *configuration.Action) error {
	path := action.Options().AsString("file", "")
	if path == "" {
		return errors.New("option 'file' is missing")
	}
	a.hookBundle.AppIO.Write("register hook fail event", true, io.VERBOSE)
	dispatcher.RegisterHookFailedSubscribers(NewCacheOnFailEventHandler(a.hookBundle, path))

	return nil
}

func NewCacheOnFail(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Action {
	a := CacheOnFail{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{info.CommitMsg}),
	}
	return &a
}

type CacheOnFailEventHandler struct {
	bundle *hooks.HookBundle
	path   string
}

func (h *CacheOnFailEventHandler) Handle(event *events.HookFailed) error {
	h.bundle.AppIO.Write("  CacheOnFail - handle failed event", true, io.VERBOSE)

	msg, err := h.bundle.Repo.CommitMessage(h.bundle.AppIO.Argument(info.ArgCommitMsgFile, ""))
	if err != nil {
		return err
	}
	if !msg.IsEmpty() {
		return os.WriteFile(h.path, []byte(msg.Message()), 0644)
	}
	return nil
}

func NewCacheOnFailEventHandler(bundle *hooks.HookBundle, path string) *CacheOnFailEventHandler {
	return &CacheOnFailEventHandler{bundle: bundle, path: path}
}
