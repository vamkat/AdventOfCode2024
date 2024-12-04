package main

import "fmt"

func main() {
	table := readFromFile()
	xmas := findXmas(table)
	mas := findMas(table)
	fmt.Println(xmas, mas)
}
