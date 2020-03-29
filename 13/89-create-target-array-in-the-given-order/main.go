package main

func main() {

}

func createTargetArray(nums []int, index []int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		num, idx := nums[i], index[i]
		var next []int
		next = append(next, res[:idx]...)
		next = append(next, num)
		next = append(next, res[idx:]...)
		res = next
	}
	return res
}
