package io

import (
	"runtime"
	"strings"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func init() {
	if runtime.GOOS == "windows" {
		DeactivateColors()
	}
}

func DeactivateColors() {
	Reset = ""
	Red = ""
	Green = ""
	Yellow = ""
	Blue = ""
	Purple = ""
	Cyan = ""
	Gray = ""
	White = ""
}

func Colorize(text string) string {
	text = strings.Replace(text, "<info>", Cyan, -1)
	text = strings.Replace(text, "</info>", Reset, -1)
	text = strings.Replace(text, "<comment>", Yellow, -1)
	text = strings.Replace(text, "</comment>", Reset, -1)
	text = strings.Replace(text, "<headline>", Purple, -1)
	text = strings.Replace(text, "</headline>", Reset, -1)
	text = strings.Replace(text, "<ok>", Green, -1)
	text = strings.Replace(text, "</ok>", Reset, -1)
	text = strings.Replace(text, "<warning>", Red, -1)
	text = strings.Replace(text, "</warning>", Reset, -1)
	return text
}
