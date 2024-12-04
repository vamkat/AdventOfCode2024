package main

import "regexp"

func findStringInstances(text string) []string {
	pattern := `mul\((\d{1,3}),(\d{1,3})\)`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllString(text, -1)
	return matches
}
