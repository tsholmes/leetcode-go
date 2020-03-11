package main

import (
	"sort"
	"strconv"
)

func main() {

}

func findRelativeRanks(nums []int) []string {
	idxs := make([]int, len(nums))
	for i := range idxs {
		idxs[i] = i
	}
	sort.Slice(idxs, func(i, j int) bool {
		return nums[idxs[i]] > nums[idxs[j]]
	})

	places := make([]int, len(nums))
	for i, n := range idxs {
		places[n] = i
	}

	res := make([]string, len(idxs))
	for i, p := range places {
		if p == 0 {
			res[i] = "Gold Medal"
		} else if p == 1 {
			res[i] = "Silver Medal"
		} else if p == 2 {
			res[i] = "Bronze Medal"
		} else {
			res[i] = strconv.Itoa(p + 1)
		}
	}
	return res
}
