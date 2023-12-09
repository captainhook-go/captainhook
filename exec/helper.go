package exec

import "strings"

// isInternalFunctionality answers if an action should trigger internal CaptainHook functionality
func isInternalFunctionality(action string) bool {
	return strings.HasPrefix(action, "CaptainHook::")
}

// splitInternalPath is determining the internal functionality to call
// Internal paths consist of two blocks seperated by .
//
// Examples:
// - CaptainHook::SOME.FUNCTIONALITY
// - CaptainHook::Branch.EnsureNaming
func splitInternalPath(action string) []string {
	actionPath := strings.Split(action, "::")[1]
	return strings.Split(actionPath, ".")
}
