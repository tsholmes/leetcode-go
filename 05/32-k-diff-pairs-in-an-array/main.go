package main

func main() {

}

func findPairs(nums []int, k int) int {
	counts := map[int]int{}
	for _, n := range nums {
		counts[n]++
	}
	res := 0
	for n := range counts {
		if k == 0 && counts[n] >= 2 {
			res++
		} else if k > 0 && counts[n+k] > 0 {
			res++
		}
	}
	return res
}
