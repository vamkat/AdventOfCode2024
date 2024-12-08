package main

import "fmt"

func main() {
	room := readFromFile()
	y, x, direction := findGuard(room)
	room, _ = moveGuard(room, y, x, direction)
	positions := countPositions(room)
	fmt.Println(positions)
	fmt.Println(checkAllObstacles(room, y, x, direction))
}
