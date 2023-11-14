package io

import (
	"bufio"
	"os"
)

type Input interface {
	Data() []string
	Argument(name, defaultValue string) string
	Arguments() map[string]string
}

type StdIn struct {
	stdInLoaded bool
	stdInData   []string
	arguments   map[string]string
}

func (s *StdIn) Data() []string {
	if !s.stdInLoaded {
		s.stdInData = s.readStdIn()
		s.stdInLoaded = true
	}
	return s.stdInData
}

func (s *StdIn) Argument(name, defaultValue string) string {
	value, ok := s.arguments[name]
	if !ok {
		value = defaultValue
	}
	return value
}

func (s *StdIn) Arguments() map[string]string {
	return s.arguments
}

func (s *StdIn) readStdIn() []string {
	var lines []string
	if s.isPiped() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			return []string{}
		}
	}
	return lines
}

func (s *StdIn) isPiped() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func NewStdIn(args map[string]string) *StdIn {
	return &StdIn{arguments: args, stdInLoaded: false}
}
