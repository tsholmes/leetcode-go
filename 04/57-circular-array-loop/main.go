package main

func main() {

}

func circularArrayLoop(nums []int) bool {
	visited := make([]bool, len(nums))
	used := make([]bool, len(nums))
	var walk func(i int, fwd bool) bool
	walk = func(i int, fwd bool) bool {
		if (nums[i] > 0) != fwd {
			return false
		}
		if used[i] {
			return true
		}
		if visited[i] {
			return false
		}
		visited[i] = true

		next := (((i + nums[i]) % len(nums)) + len(nums)) % len(nums)
		if next == i {
			return false
		}

		used[i] = true
		if walk(next, fwd) {
			return true
		}
		used[i] = false
		return false
	}
	for i := range nums {
		if walk(i, nums[i] > 0) {
			return true
		}
	}
	return false
}
