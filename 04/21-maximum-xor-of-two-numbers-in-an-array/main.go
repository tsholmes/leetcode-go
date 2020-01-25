package main

func main() {

}

func findMaximumXOR(nums []int) int {
	max := 0
	for _, n := range nums {
		for _, n2 := range nums {
			if n^n2 > max {
				max = n ^ n2
			}
		}
	}
	return max
}
