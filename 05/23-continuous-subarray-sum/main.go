package main

func main() {

}

func checkSubarraySum(nums []int, k int) bool {
	if k == 0 {
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] == 0 && nums[i+1] == 0 {
				return true
			}
		}
		return false
	}
	if k < 0 {
		k = -k
	}
	total := 0
	seen := map[int]int{}
	seen[0] = -1
	for i, n := range nums {
		total = (total + n) % k
		if j, ok := seen[total]; ok && i-j > 1 {
			return true
		} else if !ok {
			seen[total] = i
		}
	}
	return false
}
