package info

const (
	PreCommit        = "pre-commit"
	PrePush          = "pre-push"
	CommitMsg        = "commit-msg"
	PrepareCommitMsg = "prepare-commit-msg"
	PostCommit       = "post-commit"
	PostMerge        = "post-merge"
	PostCheckout     = "post-checkout"
	PostRewrite      = "post-rewrite"
	PostChange       = "post-change"
)

func GetValidHooks() []string {
	validHooks := append(GetNativeHooks(), GetVirtualHooks()...)
	return validHooks
}

func GetNativeHooks() []string {
	return []string{
		PreCommit,
		PrePush,
		CommitMsg,
		PrepareCommitMsg,
		PostCommit,
		PostMerge,
		PostCheckout,
		PostRewrite,
	}
}

func GetVirtualHooks() []string {
	return []string{
		PostChange,
	}
}

func VirtualHook(hook string) (string, bool) {
	var vHook string
	mapping := map[string]string{
		PostCheckout: PostChange,
		PostMerge:    PostChange,
		PostRewrite:  PostChange,
	}
	vHook, ok := mapping[hook]
	return vHook, ok
}
