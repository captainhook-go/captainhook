package io

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPiped() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func askForUserInput(message string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(Colorize(message))

	input, _, err := reader.ReadLine()
	return string(input), err
}

func AnswerToBool(answer string) bool {
	value := strings.ToLower(answer)
	for _, truthy := range []string{"y", "yes", "ok", "true", "+"} {
		if value == truthy {
			return true
		}
	}
	return false
}

func SplitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func SubString(input string, start int, length int) string {
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
