package main

func main() {

}

func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	counts := map[int]int{}
	for _, n := range nums1 {
		counts[n]++
	}
	for _, n := range nums2 {
		if counts[n] > 0 {
			res = append(res, n)
			counts[n] = 0
		}
	}
	return res
}
