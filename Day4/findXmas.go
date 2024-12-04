package main

func findXmas(table [][]string) int {
	var sum int

	for row := range table {
		for col := range table[row] {
			sum += checkAllDirections(table, row, col)

		}

	}
	return sum
}

func checkAllDirections(table [][]string, row, col int) int {
	var matches int
	if bottomLeftDiagonal(table, row, col, 0) {
		matches++
	}
	if reverseVertical(table, row, col, 0) {
		matches++
	}
	if topRightDiagonal(table, row, col, 0) {
		matches++
	}
	if bottomRightDiagonal(table, row, col, 0) {
		matches++
	}
	if reverseHorizontal(table, row, col, 0) {
		matches++
	}
	if horizontal(table, row, col, 0) {
		matches++
	}
	if topLeftDiagonal(table, row, col, 0) {
		matches++
	}
	if vertical(table, row, col, 0) {
		matches++
	}
	return matches
}

func topLeftDiagonal(table [][]string, row, col, letter int) bool {
	// Check in the bottom-left diagonal (row-1, col-1)
	return checkDirection(table, row, col, letter, 1, 1)
}

func vertical(table [][]string, row, col, letter int) bool {
	// Check in the bottom-left diagonal (row-1, col-1)
	return checkDirection(table, row, col, letter, 1, 0)
}

func topRightDiagonal(table [][]string, row, col, letter int) bool {
	// Check in the bottom-left diagonal (row-1, col-1)
	return checkDirection(table, row, col, letter, 1, -1)
}

func horizontal(table [][]string, row, col, letter int) bool {
	// Check in the bottom-left diagonal (row-1, col-1)
	return checkDirection(table, row, col, letter, 0, 1)
}

func reverseHorizontal(table [][]string, row, col, letter int) bool {
	// Check in the bottom-left diagonal (row-1, col-1)
	return checkDirection(table, row, col, letter, 0, -1)
}

func bottomRightDiagonal(table [][]string, row, col, letter int) bool {
	// Check in the bottom-left diagonal (row-1, col-1)
	return checkDirection(table, row, col, letter, -1, -1)
}

func bottomLeftDiagonal(table [][]string, row, col, letter int) bool {
	// Check in the bottom-left diagonal (row-1, col-1)
	return checkDirection(table, row, col, letter, -1, 1)
}

func reverseVertical(table [][]string, row, col, letter int) bool {
	// Check in the upward vertical direction (row-1, col)
	return checkDirection(table, row, col, letter, -1, 0)
}

func checkDirection(table [][]string, row, col, letter, rowDelta, colDelta int) bool {
	var letters = []string{"X", "M", "A", "S"}
	// Base case: check if out of range
	if row < 0 || col < 0 || row >= len(table) || col >= len(table[0]) {
		return false
	}

	// Check if the current letter matches the table value
	if table[row][col] == letters[letter] {
		// If it's the last letter, return true
		if letter == len(letters)-1 {
			return true
		}
		// Recurse to the next position in the direction specified by rowDelta and colDelta
		return checkDirection(table, row+rowDelta, col+colDelta, letter+1, rowDelta, colDelta)
	}
	return false
}
