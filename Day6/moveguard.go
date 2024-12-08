package main

import (
	"strconv"
	"strings"
)

func moveGuard(room [][]string, y, x int, direction string) ([][]string, bool) {
	var trackMap = make(map[[2]int]int)
	var out bool
	var loop bool
	//turn current position into an x
	room[y][x] = "X"
	//sequence of moves
	//var turningPoints [][]int
	if direction == "up" {

		for {
			room, trackMap, y, x, out, loop = moveUp(room, trackMap, y, x, out, loop)
			if out || loop {
				break
			}
			//turningPoints = append(turningPoints, []int{y, x})
			room, trackMap, y, x, out, loop = moveRight(room, trackMap, y, x, out, loop)
			if out || loop {
				break
			}
			//turningPoints = append(turningPoints, []int{y, x})
			room, trackMap, y, x, out, loop = moveDown(room, trackMap, y, x, out, loop)
			if out || loop {
				break
			}
			//turningPoints = append(turningPoints, []int{y, x})
			room, trackMap, y, x, out, loop = moveLeft(room, trackMap, y, x, out, loop)
			if out || loop {
				break
			}
			//turningPoints = append(turningPoints, []int{y, x})

		}

	}
	return room, loop
}

func moveUp(room [][]string, trackMap map[[2]int]int, y, x int, out bool, loop bool) ([][]string, map[[2]int]int, int, int, bool, bool) {

	if out || y-1 < 0 {
		return room, trackMap, y, x, true, false
	}
	if room[y-1][x] == "." || room[y-1][x] == "X" {

		room[y-1][x] = "X"
		trackMap[[2]int{y, x}] += 1000
		if strings.Contains(strconv.Itoa(trackMap[[2]int{y, x}]), "3") {
			return room, trackMap, y, x, false, true
		}
		if y-1 < 0 {
			return room, trackMap, y, x, true, false
		}

		room, trackMap, y, x, out, loop = moveUp(room, trackMap, y-1, x, false, false)

	} else {
		return room, trackMap, y, x, false, false
	}
	return room, trackMap, y, x, out, false
}

func moveDown(room [][]string, trackMap map[[2]int]int, y, x int, out, loop bool) ([][]string, map[[2]int]int, int, int, bool, bool) {
	if out || y+1 > len(room)-1 {
		return room, trackMap, y, x, true, false
	}

	if room[y+1][x] == "." || room[y+1][x] == "X" {
		room[y+1][x] = "X"
		trackMap[[2]int{y, x}] += 100
		if strings.Contains(strconv.Itoa(trackMap[[2]int{y, x}]), "3") {
			return room, trackMap, y, x, false, true
		}
		if y+1 > len(room)-1 {
			return room, trackMap, y, x, true, false
		}
		room, trackMap, y, x, out, loop = moveDown(room, trackMap, y+1, x, false, false)

	} else {
		return room, trackMap, y, x, false, false
	}
	return room, trackMap, y, x, out, loop
}

func moveRight(room [][]string, trackMap map[[2]int]int, y, x int, out, loop bool) ([][]string, map[[2]int]int, int, int, bool, bool) {
	if out || x+1 > len(room[y])-1 {
		return room, trackMap, y, x, true, false
	}
	if room[y][x+1] == "." || room[y][x+1] == "X" {
		room[y][x+1] = "X"
		trackMap[[2]int{y, x}] += 10
		if strings.Contains(strconv.Itoa(trackMap[[2]int{y, x}]), "3") {
			return room, trackMap, y, x, false, true
		}
		if x+1 > len(room[y])-1 {
			return room, trackMap, y, x, true, false
		}
		room, trackMap, y, x, out, loop = moveRight(room, trackMap, y, x+1, false, false)

	} else {
		return room, trackMap, y, x, false, false
	}
	return room, trackMap, y, x, out, loop
}

func moveLeft(room [][]string, trackMap map[[2]int]int, y, x int, out, loop bool) ([][]string, map[[2]int]int, int, int, bool, bool) {
	if out || x-1 < 0 {
		return room, trackMap, y, x, true, false
	}
	if room[y][x-1] == "." || room[y][x-1] == "X" {
		room[y][x-1] = "X"
		trackMap[[2]int{y, x}] += 1
		if strings.Contains(strconv.Itoa(trackMap[[2]int{y, x}]), "3") {
			return room, trackMap, y, x, false, true
		}
		if x-1 < 0 {
			return room, trackMap, y, x, true, false
		}
		room, trackMap, y, x, out, loop = moveLeft(room, trackMap, y, x-1, false, false)

	} else {
		return room, trackMap, y, x, false, false
	}
	return room, trackMap, y, x, out, false
}
