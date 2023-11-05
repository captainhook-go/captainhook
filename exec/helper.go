package exec

import "strings"

func isInternalFunctionality(action string) bool {
	return strings.HasPrefix(action, "CaptainHook::")
}

func splitInternalPath(action string) []string {
	actionPath := strings.Split(action, "::")[1]
	return strings.Split(actionPath, ".")
}
