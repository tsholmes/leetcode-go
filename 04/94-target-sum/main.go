package main

func main() {

}

func findTargetSumWays(nums []int, S int) int {
	memo := map[[2]int]int{}
	var search func(i int, s int) int
	search = func(i int, s int) int {
		if i == len(nums) {
			if s == 0 {
				return 1
			} else {
				return 0
			}
		}
		k := [2]int{i, s}
		if v, ok := memo[k]; ok {
			return v
		}
		res := search(i+1, s+nums[i]) + search(i+1, s-nums[i])
		memo[k] = res
		return res
	}
	return search(0, S)
}
