package util

import (
	"path"
	"strings"
)

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

func FilterByDirectory(files []string, inDir string) []string {
	var filtered []string
	for _, file := range files {
		if strings.HasPrefix(file, inDir) {
			filtered = append(filtered, file)
		}
	}
	return filtered
}

func ContainsString(haystack []string, needle string) bool {
	for _, ele := range haystack {
		if ele == needle {
			return true
		}
	}
	return false
}

func ContainsAllStrings(haystack []string, needles []string) bool {
	for _, file := range needles {
		if !ContainsString(haystack, file) {
			return false
		}
	}
	return true
}

func ContainsAnyString(haystack []string, needles []string) bool {
	for _, file := range needles {
		if ContainsString(haystack, file) {
			return true
		}
	}
	return false
}
