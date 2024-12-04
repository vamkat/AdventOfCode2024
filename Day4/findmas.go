package main

func findMas(table [][]string) int {
	var sum int

	for row := range table {
		for col := range table[row] {
			if table[row][col] == "A" {
				if checkforX(table, row, col) {
					sum++
				}

			}

		}

	}
	return sum
}

func checkforX(table [][]string, row, col int) bool {
	var diagonal1 bool
	var diagonal2 bool
	if bottomLeftDiagonalX(table, row, col, 0) || topRightDiagonalX(table, row, col, 0) {
		diagonal2 = true
	}

	if bottomRightDiagonalX(table, row, col, 0) || topLeftDiagonalX(table, row, col, 0) {
		diagonal1 = true
	}
	if diagonal1 && diagonal2 {
		return true
	}
	return false
}

func topLeftDiagonalX(table [][]string, row, col, letter int) bool {
	// Check in the bottom-left diagonal (row-1, col-1)
	return checkDirectionX(table, row-1, col-1, letter, 1, 1)
}

func topRightDiagonalX(table [][]string, row, col, letter int) bool {
	// Check in the bottom-left diagonal (row-1, col-1)
	return checkDirectionX(table, row-1, col+1, letter, 1, -1)
}

func bottomRightDiagonalX(table [][]string, row, col, letter int) bool {
	// Check in the bottom-left diagonal (row-1, col-1)
	return checkDirectionX(table, row+1, col+1, letter, -1, -1)
}

func bottomLeftDiagonalX(table [][]string, row, col, letter int) bool {
	// Check in the bottom-left diagonal (row-1, col-1)
	return checkDirectionX(table, row+1, col-1, letter, -1, 1)
}

func checkDirectionX(table [][]string, row, col, letter, rowDelta, colDelta int) bool {
	var letters = []string{"M", "A", "S"}
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
		return checkDirectionX(table, row+rowDelta, col+colDelta, letter+1, rowDelta, colDelta)
	}
	return false
}
