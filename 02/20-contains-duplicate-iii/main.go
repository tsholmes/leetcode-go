package main

func main() {

}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums) && j <= i+k; j++ {
			d := nums[i] - nums[j]
			if d <= t && d >= -t {
				return true
			}
		}
	}
	return false
}
