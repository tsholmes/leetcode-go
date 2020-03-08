package main

func main() {

}

func findSubsequences(nums []int) [][]int {
	seen := map[[16]int]bool{}
	var res [][]int
	var work []int
	for mask := 0; mask < 1<<uint(len(nums)); mask++ {
		work = work[:0]
		for i, n := range nums {
			if mask&(1<<uint(i)) != 0 {
				work = append(work, n)
			}
		}
		var k [16]int
		copy(k[:], work)
		k[len(work)] = 999
		if seen[k] {
			continue
		}
		seen[k] = true
		if len(work) < 2 {
			continue
		}
		valid := true
		for i := 1; i < len(work); i++ {
			valid = valid && work[i-1] <= work[i]
		}
		if valid {
			res = append(res, append([]int{}, work...))
		}
	}
	return res
}
