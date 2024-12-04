package main

import "regexp"

func findAllPatterns(text string) []string {
	pattern := `mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllString(text, -1)
	return matches
}
