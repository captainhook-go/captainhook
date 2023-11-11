package io

import (
	"runtime"
	"strings"
)

var Reset string
var Red string
var Green string
var Yellow string
var Blue string
var Purple string
var Cyan string
var Gray string
var White string

func init() {
	if runtime.GOOS == "windows" {
		DeactivateColors()
	} else {
		ActivateColors()
	}
}

func ColorStatus(on bool) {
	if on {
		ActivateColors()
	} else {
		DeactivateColors()
	}
}

func ActivateColors() {
	Reset = "\033[0m"
	Red = "\033[31m"
	Green = "\033[32m"
	Yellow = "\033[33m"
	Blue = "\033[34m"
	Purple = "\033[35m"
	Cyan = "\033[36m"
	Gray = "\033[37m"
	White = "\033[97m"
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
	text = strings.Replace(text, "<strong>", Blue, -1)
	text = strings.Replace(text, "</strong>", Reset, -1)
	return text
}
