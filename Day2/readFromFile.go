package main

import (
	"os"
	"strings"
)

func readFromFile() []string {

	f, _ := os.ReadFile("input.txt")
	input := strings.Split(string(f), "\r\n")
	return input
}
