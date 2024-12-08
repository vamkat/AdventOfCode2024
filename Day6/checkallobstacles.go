package main

//make a copy of the solved room and try an obstacle on every X (guard's path)
func checkAllObstacles(room [][]string, startY, startX int, direction string) int {
	var obstacles int

	for row := range room {
		for col := range room {

			if room[row][col] == "X" {
				if !(row == startY && col == startX) {
					copyRoom := deepCopy(room)
					copyRoom[row][col] = "#"
					_, loop := moveGuard(copyRoom, startY, startX, direction)
					if loop {
						obstacles++
					}
				}
			}
		}
	}
	return obstacles
}

func deepCopy(slice [][]string) [][]string {
	newSlice := make([][]string, len(slice))
	for i := range slice {
		newSlice[i] = make([]string, len(slice[i]))
		copy(newSlice[i], slice[i])
	}
	return newSlice
}
