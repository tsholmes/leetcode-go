package main

func main() {

}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nexts := map[int]int{}
	for i, n := range nums2 {
		next := -1
		for _, n2 := range nums2[i:] {
			if n2 > n {
				next = n2
				break
			}
		}
		nexts[n] = next
	}
	res := make([]int, len(nums1))
	for i, n := range nums1 {
		res[i] = nexts[n]
	}
	return res
}
