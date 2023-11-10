package io

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getUserInput(message string) (string, error) {
	fmt.Print(Colorize(message))

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return "", err
	}
	return input, nil
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
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	stats, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %s", path)
	}
	if stats.IsDir() {
		return nil, fmt.Errorf("given file path is a directory: %s", path)
	}

	closeErr := file.Close()
	if closeErr != nil {
		return nil, fmt.Errorf("error closing file handle for: %s", path)
	}

	data, readErr := os.ReadFile(path)
	if readErr != nil {
		return nil, fmt.Errorf("could not read configuration file at: %s", path)
	}
	return data, nil
}
