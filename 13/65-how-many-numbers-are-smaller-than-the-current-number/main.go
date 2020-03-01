package main

func main() {

}

func smallerNumbersThanCurrent(nums []int) []int {
	res := make([]int, len(nums))
	for i, n1 := range nums {
		for _, n2 := range nums {
			if n2 < n1 {
				res[i]++
			}
		}
	}
	return res
}
