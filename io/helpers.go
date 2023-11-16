package io

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func AnswerToBool(answer string) bool {
	value := strings.ToLower(answer)
	truth := []string{"y", "yes", "ok", "true", "+", "1"}

	return slices.Contains(truth, value)
}

func SplitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func PrefixLinesInString(s, prefix string) string {
	return strings.Join(
		PrefixLines(
			SplitLines(s),
			prefix,
		),
		"\n",
	)
}

func PrefixLines(lines []string, prefix string) []string {
	var prefixed []string
	for _, s := range lines {
		prefixed = append(prefixed, prefix+s)
	}
	return prefixed
}

func SubString(input string, start int, length int) string {
	if input == "" {
		return ""
	}

	asRunes := []rune(input)

	if length == 0 {
		length = len(asRunes)
	}
	if start < 0 {
		start = len(asRunes) + start
	}
	if length < 0 {
		length = len(asRunes) + length
	}
	if start >= len(asRunes) {
		return ""
	}
	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}
	return string(asRunes[start : start+length])
}

func FileExists(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		return false
	}
	_, err = file.Stat()
	if err != nil {
		return false
	}
	return true
}

func ReadFile(path string) ([]byte, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if fileInfo.IsDir() {
		return nil, fmt.Errorf("given file path is a directory: %s", path)
	}

	data, readErr := os.ReadFile(path)
	if readErr != nil {
		return nil, fmt.Errorf("could not read configuration file at: %s", path)
	}
	return data, nil
}
