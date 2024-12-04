package main

import "strings"

func getNumbers(matches []string) [][]string {
	var numbers [][]string
	for _, entry := range matches {
		entry = strings.TrimLeft(entry, "mul(")
		entry = strings.TrimRight(entry, ")")
		numbersStr := strings.Split(entry, ",")
		numbers = append(numbers, numbersStr)
	}
	return numbers
}
