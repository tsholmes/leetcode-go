package main

import "sort"

func main() {

}

func topKFrequent(nums []int, k int) []int {
	counts := map[int]int{}
	for _, n := range nums {
		counts[n]++
	}
	var res []int
	for k := range counts {
		res = append(res, k)
	}
	sort.Slice(res, func(i, j int) bool {
		return counts[res[i]] > counts[res[j]]
	})
	if len(res) > k {
		res = res[:k]
	}
	return res
}
