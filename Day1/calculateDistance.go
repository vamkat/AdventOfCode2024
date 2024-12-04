package main

import "slices"

func calculateDistance(list1, list2 []int) int {

	//sort lists
	slices.Sort(list1)
	slices.Sort(list2)

	var distances []int
	for i := range list1 {
		distance := list1[i] - list2[i]
		if distance < 0 {
			distance = distance * (-1)
		}
		distances = append(distances, distance)
	}

	//add up distances
	var sum int
	for _, distance := range distances {
		sum += distance
	}
	return sum
}
