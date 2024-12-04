package main

import "strconv"

func getResult(numbers [][]string) int {
	var sum int
	for _, nums := range numbers {
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		sum += num1 * num2
	}
	return sum
}
