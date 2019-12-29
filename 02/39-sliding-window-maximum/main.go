package main

import "fmt"

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4))
	fmt.Println(maxSlidingWindow([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 4))
	fmt.Println(maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3))
	fmt.Println(maxSlidingWindow([]int{2, 0, 1, 0, 2}, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	starts := make([]int, len(nums))
	ends := make([]int, len(nums))

	for i := range nums {
		if i%k == 0 {
			starts[i] = nums[i]
		} else {
			starts[i] = max(starts[i-1], nums[i])
		}

		b := k
		if (i-i%k)+k > len(nums) {
			b = len(nums) % k
		}
		j := i - (i % k) + (b - (i % k) - 1)
		if i%k == 0 {
			ends[j] = nums[j]
		} else {
			ends[j] = max(ends[j+1], nums[j])
		}
	}

	res := make([]int, len(nums)-k+1)
	for i := range res {
		j := i + k - 1
		res[i] = max(ends[i], starts[j])
	}

	return res
}
