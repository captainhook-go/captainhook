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

// GetValidHooks is returning all hooks supported by CaptainHook git native and virtual ones.
func GetValidHooks() []string {
	validHooks := append(GetNativeHooks(), GetVirtualHooks()...)
	return validHooks
}

// GetNativeHooks is returning all hook native to git
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

// GetVirtualHooks is retuning all virtual hooks provided by CaptainHook
func GetVirtualHooks() []string {
	return []string{
		PostChange,
	}
}

// VirtualHook returns the virtual hook a native hook triggers.
// Examples:
// - post-checkout triggers post-change
// - post-merge triggers post-change
// ...
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
