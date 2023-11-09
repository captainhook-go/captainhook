package io

import (
	"bufio"
	"fmt"
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
