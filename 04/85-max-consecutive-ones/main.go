package main

func main() {

}

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	cur := 0
	for len(nums) > 0 {
		if nums[0] == 1 {
			cur++
			if cur > max {
				max = cur
			}
		} else {
			cur = 0
		}
		nums = nums[1:]
	}
	return max
}
