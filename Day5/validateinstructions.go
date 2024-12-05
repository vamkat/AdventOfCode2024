package main

func findValid(instructionsList, rules [][]int) (int, [][]int) {
	var sum int
	var instructionsToFix [][]int
	for _, instructions := range instructionsList {
		indexMap := makeIndexMap(instructions)
		if checkValid(indexMap, rules) {
			//find median
			median := instructions[(len(instructions)-1)/2]
			sum += median
		} else {
			instructionsToFix = append(instructionsToFix, instructions)

		}

	}

	return sum, instructionsToFix
}

func makeIndexMap(instructions []int) map[int]int {
	//for every instruction set make map of indices
	indexMap := make(map[int]int)
	for i, number := range instructions {
		indexMap[number] = i
	}

	return indexMap
}

func checkValid(indexMap map[int]int, rules [][]int) bool {

	for _, orderPair := range rules {
		index1, ok1 := indexMap[orderPair[0]]
		index2, ok2 := indexMap[orderPair[1]]
		if ok1 && ok2 {
			if index1 > index2 {
				return false
			}
		}
	}

	return true
}

func fixInstructions(invalidInstructionsList [][]int, rules [][]int) int {
	var sum int
	for _, instructions := range invalidInstructionsList {
		fixedInstructions := fixOrder(instructions, rules)
		median := fixedInstructions[(len(fixedInstructions)-1)/2]
		sum += median
	}
	return sum
}

func fixOrder(instructions []int, rules [][]int) []int {
	//make map of indices
	indexMap := makeIndexMap(instructions)
	var swapped bool

	for {
		if checkValid(indexMap, rules) {
			return instructions
		}
		swapped = false
		for _, orderPair := range rules {
			index1, ok1 := indexMap[orderPair[0]]
			index2, ok2 := indexMap[orderPair[1]]
			if ok1 && ok2 && index1 > index2 {

				instructions[index1], instructions[index2] = instructions[index2], instructions[index1]
				swapped = true
				//update map
				indexMap[orderPair[0]] = index2
				indexMap[orderPair[1]] = index1
			}
		}

		if !swapped {
			break
		}
	}
	return instructions

}
