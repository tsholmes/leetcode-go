package main

func main() {

}

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	copy(res, nums)
	for i := len(nums) - 2; i >= 0; i-- {
		res[i] = nums[i] * res[i+1]
	}
	for i := 1; i < len(nums); i++ {
		nums[i] *= nums[i-1]
	}

	for i := range res {
		prod := 1
		if i > 0 {
			prod *= nums[i-1]
		}
		if i < len(nums)-1 {
			prod *= res[i+1]
		}
		res[i] = prod
	}

	return res
}
