package main

import (
	"strconv"
	"strings"
)

func countSafe(input []string) (int, int) {
	var sum int
	var extras int
	for i := range input {
		var numbers []int
		numbersS := strings.Split(input[i], " ")
		for j := range numbersS {
			num, _ := strconv.Atoi(numbersS[j])
			numbers = append(numbers, num)
		}
		if checkSafe(numbers) {
			sum++
			//fmt.Println(input[i])
		} else {
			extras += countExtra(numbers)
		}
	}
	return sum, extras
}
