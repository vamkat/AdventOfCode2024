package main

func checkSafe(numbers []int) bool {
	var sorted bool
	if numbers[0] > numbers[1] {
		sorted = true
		for n := 0; n < len(numbers)-1; n++ {
			diff := numbers[n] - numbers[n+1]
			if diff <= 3 && diff >= 1 {
				//all good
			} else {
				sorted = false
				break
			}
		}

	} else if numbers[0] < numbers[1] {
		sorted = true
		for n := 0; n < len(numbers)-1; n++ {
			diff := numbers[n+1] - numbers[n]
			if diff <= 3 && diff >= 1 {
				//all good
			} else {
				sorted = false
				break
			}

		}

	}
	return sorted
}
