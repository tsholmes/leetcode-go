package main

func main() {

}

func findMaxConsecutiveOnes(nums []int) int {
	prev := -1
	cur := 0
	max := 0
	for _, n := range nums {
		if n == 1 {
			cur++
		} else {
			prev = cur
			cur = 0
		}
		if prev+cur+1 > max {
			max = prev + cur + 1
		}
	}
	return max
}
