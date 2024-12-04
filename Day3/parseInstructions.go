package main

func parseInstructions(instructions []string) []string {
	keep := true
	var validInstructions []string
	for _, entry := range instructions {
		if entry == "do()" {
			keep = true
		} else if entry == "don't()" {
			keep = false
		} else if keep {
			validInstructions = append(validInstructions, entry)
		}
	}
	return validInstructions
}
