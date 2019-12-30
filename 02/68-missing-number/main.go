package main

func main() {

}

func missingNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == i || nums[i] == len(nums) {
			continue
		}
		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		i--
	}
	for i := range nums {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}
