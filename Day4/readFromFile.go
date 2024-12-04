package main

import (
	"os"
	"strings"
)

func readFromFile() [][]string {

	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(f), "\r\n")
	var table [][]string
	for _, line := range lines {
		row := strings.Split(line, "")
		table = append(table, row)
	}
	return table
}
