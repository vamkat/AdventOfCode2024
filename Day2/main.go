package main

import (
	"fmt"
)

func main() {

	input := readFromFile()
	count1, extras := countSafe(input)

	fmt.Println(count1, count1+extras)
}
