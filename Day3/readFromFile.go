package main

import (
	"os"
)

func readFromFile() string {

	f, _ := os.ReadFile("input.txt")
	return string(f)
}
