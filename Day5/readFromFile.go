package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func readFromFile() ([][]int, [][]int) {
	//read input
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("can't read file")
	}
	lines := strings.Split(string(f), "\r\n")
	var instructionsList [][]int
	for _, line := range lines {
		var instructions []int
		instructionsS := strings.Split(line, ",")
		for _, number := range instructionsS {
			num, _ := strconv.Atoi(number)
			instructions = append(instructions, num)
		}
		instructionsList = append(instructionsList, instructions)
	}
	r, err := os.ReadFile("rules.txt")
	if err != nil {
		log.Fatal("can't read file")
	}
	rules := strings.Split(string(r), "\r\n")
	var orderRules [][]int
	for _, line := range rules {
		order := strings.Split(line, "|")
		var orderedNums []int
		for _, number := range order {
			num, _ := strconv.Atoi(number)
			orderedNums = append(orderedNums, num)
		}
		orderRules = append(orderRules, orderedNums)
	}

	return instructionsList, orderRules
}
