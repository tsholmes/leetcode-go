package main

func main() {

}

func countSmaller(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	res := make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				res[i]++
			}
		}
	}
	return res
}
