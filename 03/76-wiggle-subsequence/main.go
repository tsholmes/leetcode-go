package main

func main() {

}

func wiggleMaxLength(nums []int) int {
	N := len(nums)
	up := make([]int, N)
	down := make([]int, N)
	max := 0
	for i, n := range nums {
		up[i] = 1
		down[i] = 1
		for j, n2 := range nums[:i] {
			if n > n2 && up[i] <= down[j] {
				up[i] = down[j] + 1
			}
			if n < n2 && down[i] <= up[j] {
				down[i] = up[j] + 1
			}
		}
		if up[i] > max {
			max = up[i]
		}
		if down[i] > max {
			max = down[i]
		}
	}
	return max
}
