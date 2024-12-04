package main

func similarityScore(list1, list2 []int) int {
	entries := make(map[int]int)

	//add to map if number in list1 is found in list2
	for _, num := range list1 {
		for _, entry := range list2 {
			if num == entry {
				entries[num]++
			}
		}
	}
	var score int
	for num, frequency := range entries {
		score += num * frequency
	}
	return score
}
