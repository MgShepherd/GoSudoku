package main

func sliceFromRange(start, end int) []int {
	if end - start <= 0 {
		return []int{}
	}
	rSlice := make([]int, end-start+1)

	for i := 0; i <= end-start; i++ {
		rSlice[i] = start+i
	}
	return rSlice
}
