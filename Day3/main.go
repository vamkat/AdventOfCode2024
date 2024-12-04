package main

import "fmt"

func main() {
	text := readFromFile()

	//part1
	matches := findStringInstances(text)
	numbers := getNumbers(matches)
	fmt.Println(getResult(numbers))

	//part2
	fullMatches := findAllPatterns(text)
	valid := parseInstructions(fullMatches)
	nums := getNumbers(valid)
	fmt.Println(getResult(nums))

}
