package main

func main() {

}

func subarraySum(nums []int, k int) int {
	tsum := 0
	seen := map[int]int{0: 1}
	res := 0
	for _, n := range nums {
		tsum += n
		res += seen[tsum-k]
		seen[tsum]++
	}
	return res
}
