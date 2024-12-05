package main

import "fmt"

func main() {
	instructions, rules := readFromFile()
	//parseOrder(rules)
	sum, instructionsToFix := findValid(instructions, rules)
	sum2 := fixInstructions(instructionsToFix, rules)
	fmt.Println(sum, sum2)
}
