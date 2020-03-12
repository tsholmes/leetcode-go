package main

func main() {

}

func findMaxLength(nums []int) int {
	diff := 0
	seen := map[int]int{}
	seen[0] = -1
	max := 0
	for i, n := range nums {
		if n == 1 {
			diff++
		} else {
			diff--
		}
		if v, ok := seen[diff]; ok && i-v > max {
			max = i - v
		} else if !ok {
			seen[diff] = i
		}
	}
	return max
}
