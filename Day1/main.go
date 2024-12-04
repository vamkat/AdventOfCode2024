package main

import (
	"fmt"
)

func main() {
	list1, list2 := readFromFile()
	//for part1:
	fmt.Println(calculateDistance(list1, list2))
	//for part2:
	fmt.Println(similarityScore(list1, list2))
}
