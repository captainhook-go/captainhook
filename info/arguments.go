package info

// Names of all hook arguments to map them from position to name
const (
	ArgCommitMsgFile = "message-file"
	ArgGitCommand    = "git-command"
	ArgHash          = "hash"
	ArgPreviousHead  = "previous-hash"
	ArgMode          = "mode"
	ArgNewHead       = "new-head"
	ArgSquash        = "squash"
	ArgTarget        = "target"
	ArgURL           = "url"
	ArgCommand       = "command"
)

var (
	HookArgs = map[string][]string{
		CommitMsg:        {ArgCommitMsgFile},
		PostCheckout:     {ArgPreviousHead, ArgNewHead, ArgMode},
		PostCommit:       {},
		PostMerge:        {ArgSquash},
		PostRewrite:      {ArgGitCommand},
		PreCommit:        {},
		PrePush:          {ArgTarget, ArgURL},
		PrepareCommitMsg: {ArgCommitMsgFile, ArgMode, ArgHash},
	}
)

// HookArguments returns the arguments for a given hook
func HookArguments(hook string) []string {
	args, ok := HookArgs[hook]
	if !ok {
		return []string{}
	}
	return args
}

// AllHookArguments returns a slice of all available arguments
// It is used to detect if a placeholder config is correct
func AllHookArguments() []string {
	return []string{
		ArgCommitMsgFile,
		ArgMode,
		ArgHash,
		ArgTarget,
		ArgURL,
		ArgPreviousHead,
		ArgNewHead,
		ArgSquash,
	}
}
