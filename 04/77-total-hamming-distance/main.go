package main

func main() {

}

func totalHammingDistance(nums []int) int {
	var ones [32]int
	var zeros [32]int
	for _, n := range nums {
		for i := uint(0); i < 32; i++ {
			if n&(1<<i) != 0 {
				ones[i]++
			} else {
				zeros[i]++
			}
		}
	}
	var res int
	for i := 0; i < 32; i++ {
		res += ones[i] * zeros[i]
	}
	return res
}
