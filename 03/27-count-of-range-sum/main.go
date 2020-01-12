package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(countRangeSum([]int{-2, 5, -1}, -2, 2))
}

func countRangeSum(nums []int, lower int, upper int) int {
	sumis := map[int][]int{}
	csum := 0
	for i := range nums {
		csum += nums[i]
		sumis[csum] = append(sumis[csum], i)
	}
	var sums []int
	for k := range sumis {
		sums = append(sums, k)
	}
	sort.Ints(sums)

	res := 0
	csum = 0
	for i := range nums {
		csum += nums[i]
		if csum >= lower && csum <= upper {
			res++
		}
		si := sort.Search(len(sums), func(j int) bool {
			return csum-sums[j] <= upper
		})
		for si < len(sums) && csum-sums[si] >= lower {
			sis := sumis[sums[si]]
			ci := sort.Search(len(sis), func(j int) bool { return sis[j] >= i })
			res += ci
			si++
		}
	}
	return res
}
