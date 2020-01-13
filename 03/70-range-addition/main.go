package main

import "sort"

func main() {

}

func getModifiedArray(length int, updates [][]int) []int {
	var starts [][2]int
	var ends [][2]int
	for _, u := range updates {
		starts = append(starts, [2]int{u[0], u[2]})
		ends = append(ends, [2]int{u[1], u[2]})
	}
	sort.Slice(starts, func(i, j int) bool {
		return starts[i][0] < starts[j][0]
	})
	sort.Slice(ends, func(i, j int) bool {
		return ends[i][0] < ends[j][0]
	})

	val := 0
	res := make([]int, length)
	for i := range res {
		for len(starts) > 0 && starts[0][0] <= i {
			val += starts[0][1]
			starts = starts[1:]
		}
		res[i] = val
		for len(ends) > 0 && ends[0][0] <= i {
			val -= ends[0][1]
			ends = ends[1:]
		}
	}
	return res
}
