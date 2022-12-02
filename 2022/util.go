package main

func sumInts(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}
