package main

func main() {

}

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	mins := make([]int, len(nums))
	mins[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < mins[i-1] {
			mins[i] = nums[i]
		} else {
			mins[i] = mins[i-1]
		}
	}
	var stack []int
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == mins[i] {
			continue
		}
		for len(stack) > 0 && stack[len(stack)-1] <= mins[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 && stack[len(stack)-1] < nums[i] {
			return true
		}
		stack = append(stack, nums[i])
	}

	return false
}
