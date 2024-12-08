package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	results, numbers := readFromFile()
	sum := solve(results, numbers)
	fmt.Println(sum)
}

func solve(results []int, numbers [][]int) int {
	var sum int
	for i, factors := range numbers {
		sum += solveForThisNumberSet(results[i], factors)
	}
	return sum
}

func solveForThisNumberSet(wantedResult int, numbers []int) int {
	//var sum int
	possibleCombinations := generateCombinations(len(numbers) - 1)
	for _, operators := range possibleCombinations {
		//need to do multiplications first
		if solveForThisOperatorSet(wantedResult, numbers, operators) > 0 {
			return wantedResult
		}
	}

	return 0
}

func solveForThisOperatorSet(wantedResult int, numbers []int, operators []string) int {
	var result int
	result += numbers[0]
	for i, operator := range operators {

		if operator == "+" {
			result = result + numbers[i+1]
		} else if operator == "*" {

			result = result * numbers[i+1]
		} else if operator == "||" {
			numString := strconv.Itoa(result) + strconv.Itoa(numbers[i+1])
			result, _ = strconv.Atoi(numString)
		}
	}
	if result == wantedResult {
		return wantedResult
	} else {
		return 0
	}

}

func generateCombinations(length int) [][]string {
	var result [][]string
	totalCombinations := int(math.Pow(3, float64(length))) // Calculate 3^length

	// Loop through numbers from 0 to totalCombinations - 1
	for i := 0; i < totalCombinations; i++ {
		var combination []string
		for j := 0; j < length; j++ {
			// Determine the character by taking the remainder when dividing by 3
			switch (i / int(math.Pow(3, float64(length-j-1)))) % 3 {
			case 0:
				combination = append(combination, "*")
			case 1:
				combination = append(combination, "+")
			case 2:
				combination = append(combination, "||")
			}
		}
		result = append(result, combination)
	}

	return result
}
