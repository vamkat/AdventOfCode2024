package main

//returns guard's initial position
//takes into account all 4 directions (not needed)
func findGuard(room [][]string) (int, int, string) {
	//assuming one guard: <>^v
	for row := range room {
		for col := range room {
			if room[row][col] != "." && room[row][col] != "#" {
				switch room[row][col] {
				case "^":
					return row, col, "up"
				case "v":
					return row, col, "down"
				case "<":
					return row, col, "left"
				case ">":
					return row, col, "right"
				default:
					return row, col, ""
				}

			}
		}
	}
	return 0, 0, ""
}
