package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func readFromFile() ([]int, []int) {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("can't read file")
	}
	fString := strings.ReplaceAll(string(f), "\r\n", "   ")
	locations := strings.Split(fString, "   ")

	var list1 []int
	var list2 []int

	for i := range locations {
		num, _ := strconv.Atoi(locations[i])
		if i%2 == 0 {
			//odd because list starts with 0

			list1 = append(list1, num)
		} else {
			list2 = append(list2, num)
		}
	}
	return list1, list2
}
