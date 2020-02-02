package main

func main() {

}

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if n != i+1 && nums[n-1] != n {
			nums[i], nums[n-1] = nums[n-1], nums[i]
			i--
		}
	}
	count := 0
	for i, n := range nums {
		if n != i+1 {
			nums[count] = i + 1
			count++
		}
	}
	nums = nums[:count]
	return nums
}
