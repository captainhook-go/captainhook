package file

import (
	"fmt"
	"strings"
)

const (
	Quote    = "(\"|')"
	OptQuote = Quote + "?"
	Connect  = "\\s*(:|=>|=|:=)\\s*"
)

var presets = map[string]func() []string{
	"aws":    AwsPatterns,
	"github": GitHubPatterns,
	"google": GooglePatterns,
	"stripe": StripePatterns,
}

func GetPreset(name string) ([]string, error) {
	presetFunc, ok := presets[strings.ToLower(name)]
	if !ok {
		return []string{}, fmt.Errorf("preset %s not found", name)
	}
	return presetFunc(), nil
}

func AwsPatterns() []string {
	aws := "(AWS|aws|Aws)?_?"
	return []string{
		// AWS token
		"(A3T[A-Z0-9]|AKIA|AGPA|AIDA|AROA|AIPA|ANPA|ANVA|ASIA)[A-Z0-9]{16}",
		// AWS secrets, keys, access token
		OptQuote + aws + "(SECRET|secret|Secret)?_?(ACCESS|access|Access)?_?(KEY|key|Key)" + OptQuote + Connect + OptQuote + "[A-Za-z0-9/\\+=]{40}" + OptQuote,
		// AWS account id
		OptQuote + aws + "(ACCOUNT|account|Account)_?(ID|id|Id)?" + OptQuote + Connect + OptQuote + "[0-9]{4}\\-?[0-9]{4}\\-?[0-9]{4}" + OptQuote,
	}
}

func GitHubPatterns() []string {
	return []string{
		// Personal Access Token (Classic)
		OptQuote + "(ghp_[a-zA-Z0-9]{36})" + OptQuote,
		// Personal Access Token (Fine-Grained)
		OptQuote + "(github_pat_[a-zA-Z0-9]{22}_[a-zA-Z0-9]{59})" + OptQuote,
		// User-To-Server Access Token
		OptQuote + "(ghu_[a-zA-Z0-9]{36})" + OptQuote,
		// Server-To-Server Access Token
		OptQuote + "(ghs_[a-zA-Z0-9]{36})" + OptQuote,
	}
}

func StripePatterns() []string {
	return []string{
		// Standard API Key & Restricted API Key
		OptQuote + "(sk_live_[0-9a-z]{24})" + OptQuote,
	}
}

func GooglePatterns() []string {
	return []string{
		OptQuote + "(AIza[0-9A-Za-z\\-_]{35})" + OptQuote,
	}
}
