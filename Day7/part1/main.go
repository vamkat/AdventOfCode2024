package main

import "fmt"

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
		}
	}
	if result == wantedResult {
		return wantedResult
	} else {
		return 0
	}

	// var numbersToAdd []int
	// var product int
	// var multiplying bool
	// for i, operator := range operators {

	// 	if operator == "+" {
	// 		if multiplying {
	// 			numbersToAdd = append(numbersToAdd, product)
	// 			product = 0
	// 			multiplying = false
	// 		} else {
	// 			numbersToAdd = append(numbersToAdd, numbers[i])
	// 		}
	// 		if i == len(operators)-1 {
	// 			numbersToAdd = append(numbersToAdd, numbers[i+1])
	// 		}

	// 	} else if operator == "*" {

	// 		if !multiplying {
	// 			product = numbers[i] * numbers[i+1]
	// 			multiplying = true
	// 		} else {
	// 			product = product * numbers[i] * numbers[i+1]
	// 		}
	// 		if i == len(operators)-1 {
	// 			numbersToAdd = append(numbersToAdd, product)
	// 		}

	// 	}
	// }
	// //now add everything up
	// var sum int
	// for _, number := range numbersToAdd {
	// 	sum += number
	// }
	// if sum == wantedResult {
	// 	return wantedResult
	// } else {
	// 	return 0
	// }
}

func generateCombinations(length int) [][]string {
	var result [][]string
	totalCombinations := 1 << length // 2^length

	// Loop through numbers from 0 to totalCombinations - 1
	for i := 0; i < totalCombinations; i++ {
		var combination []string
		for j := 0; j < length; j++ {
			// If the j-th bit is 1, append 'B', otherwise append 'A'
			if (i & (1 << (length - j - 1))) != 0 {
				combination = append(combination, "*")
			} else {
				combination = append(combination, "+")
			}
		}
		result = append(result, combination)
	}

	return result
}
