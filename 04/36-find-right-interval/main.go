package main

func main() {

}

func findRightInterval(intervals [][]int) []int {
	res := make([]int, len(intervals))
	for i, iv1 := range intervals {
		res[i] = -1
		for j, iv2 := range intervals {
			if iv2[0] >= iv1[1] {
				if res[i] == -1 || iv2[0] < intervals[res[i]][0] {
					res[i] = j
				}
			}
		}
	}
	return res
}
