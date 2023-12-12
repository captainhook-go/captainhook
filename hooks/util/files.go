package util

import (
	"path"
	"slices"
	"strings"
)

// FilterByType removes all files from a list that do not have the suffix ofType
// The ofType argument should include the dot ".txt", ".go" ...
func FilterByType(files []string, ofType string) []string {
	var filtered []string
	for _, file := range files {
		ext := path.Ext(file)
		if ext == ofType {
			filtered = append(filtered, file)
		}
	}
	return filtered
}

// FilterByDirectory removes all files from a slice that are not in a given directory
func FilterByDirectory(files []string, inDir string) []string {
	var filtered []string
	for _, file := range files {
		if strings.HasPrefix(file, inDir) {
			filtered = append(filtered, file)
		}
	}
	return filtered
}

// ContainsAllStrings checks if a haystack contains all needles
func ContainsAllStrings(haystack []string, needles []string) bool {
	for _, file := range needles {
		if !slices.Contains(haystack, file) {
			return false
		}
	}
	return true
}

// ContainsAnyString checks if a haystack contains any needle
func ContainsAnyString(haystack []string, needles []string) bool {
	for _, file := range needles {
		if slices.Contains(haystack, file) {
			return true
		}
	}
	return false
}
