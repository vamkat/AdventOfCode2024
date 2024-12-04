package main

func countExtra(numbers []int) int {

	for i := 0; i < len(numbers); i++ {
		var numbersTrunc []int
		for j, n := range numbers {
			if j != i {

				numbersTrunc = append(numbersTrunc, n)
			}

		}
		if checkSafe(numbersTrunc) {
			return 1
		}

	}
	return 0
}
