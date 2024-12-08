package main

//checks the whole grid for Xs
func countPositions(room [][]string) int {
	var positions int
	for row := range room {
		for col := range room {
			if room[row][col] != "#" && room[row][col] != "." {
				positions++

			}
		}
	}
	return positions
}
