package info

const (
	PRE_COMMIT         = "pre-commit"
	PRE_PUSH           = "pre-push"
	COMMIT_MSG         = "commit-msg"
	PREPARE_COMMIT_MSG = "prepare-commit-msg"
	POST_COMMIT        = "post-commit"
	POST_MERGE         = "post-merge"
	POST_CHECKOUT      = "post-checkout"
	POST_REWRITE       = "post-rewrite"
	POST_CHANGE        = "post-change"
)

func GetValidHooks() []string {
	validHooks := append(GetNativeHooks(), GetVirtualHooks()...)
	return validHooks
}

func GetNativeHooks() []string {
	return []string{
		PRE_COMMIT,
		PRE_PUSH,
		COMMIT_MSG,
		PREPARE_COMMIT_MSG,
		POST_COMMIT,
		POST_MERGE,
		POST_CHECKOUT,
		POST_REWRITE,
	}
}

func GetVirtualHooks() []string {
	return []string{
		POST_CHANGE,
	}
}

func VirtualHook(hook string) (string, bool) {
	var vHook string
	mapping := map[string]string{
		POST_CHECKOUT: POST_CHANGE,
		POST_MERGE:    POST_CHANGE,
		POST_REWRITE:  POST_CHANGE,
	}
	vHook, ok := mapping[hook]
	return vHook, ok
}
