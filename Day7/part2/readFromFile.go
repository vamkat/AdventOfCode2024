package main

import (
	"os"
	"strconv"
	"strings"
)

func readFromFile() ([]int, [][]int) {
	var results []int
	var numbers [][]int

	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(f), "\r\n")
	for _, line := range lines {
		tmp := strings.Split(line, ":")
		result, _ := strconv.Atoi(tmp[0])
		results = append(results, result)
		factors := strings.Split(strings.Trim(tmp[1], " "), " ")
		var currentNums []int
		for _, factor := range factors {

			number, _ := strconv.Atoi(factor)
			currentNums = append(currentNums, number)
		}
		numbers = append(numbers, currentNums)

	}
	return results, numbers
}
