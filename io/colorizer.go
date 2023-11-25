package io

import (
	"runtime"
	"strings"
)

var Color = NewColorizer()

func ColorStatus(on bool) {
	if on {
		Color.activate()
	} else {
		Color.deactivate()
	}
}

func Colorize(text string) string {
	return Color.Colorize(text)
}

// Colorizer adds ascii color to a string
type Colorizer struct {
	canColorize bool
	status      bool
	colors      map[string]string
	tags        map[string]string
}

// Colorize only adds color to a string if the terminal supports it
func (c *Colorizer) Colorize(text string) string {
	if c.status && c.canColorize {
		return c.colorizeText(text)
	}
	return c.removeTags(text)
}

// activate activates display of colors
func (c *Colorizer) activate() {
	c.status = true
}

// deactivate deactivates display of colors
func (c *Colorizer) deactivate() {
	c.status = false
}

// colorizeText replaces <tag> with ascii color codes
func (c *Colorizer) colorizeText(text string) string {
	for tag, color := range c.tags {
		text = strings.Replace(text, "<"+tag+">", c.colors[color], -1)
		text = strings.Replace(text, "</"+tag+">", c.colors["Reset"], -1)
	}
	return text
}

// removeTags removes all known tags from a string
func (c *Colorizer) removeTags(text string) string {
	for tag := range c.tags {
		text = strings.Replace(text, "<"+tag+">", "", -1)
		text = strings.Replace(text, "</"+tag+">", "", -1)
	}
	return text
}

// NewColorizer creates a new Colorizer and checks if the current terminal supports colors
// Currently windows is just a hard disable
func NewColorizer() *Colorizer {
	isAbleToColorize := true
	if runtime.GOOS == "windows" {
		isAbleToColorize = false
	}
	return &Colorizer{
		isAbleToColorize,
		true,
		map[string]string{
			"Reset":  "\033[0m",
			"Red":    "\033[31m",
			"Green":  "\033[32m",
			"Yellow": "\033[33m",
			"Blue":   "\033[34m",
			"Purple": "\033[35m",
			"Cyan":   "\033[36m",
			"Gray":   "\033[37m",
			"White":  "\033[97m",
		},
		map[string]string{
			"info":     "Cyan",
			"comment":  "Yellow",
			"headline": "Purple",
			"ok":       "Green",
			"warning":  "red",
			"strong":   "blue",
		},
	}
}
