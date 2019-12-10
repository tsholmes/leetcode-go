package main

import (
	"fmt"
	"math/bits"
	"sort"
)

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
	fmt.Println(permuteUnique([]int{1, 1, 1, 1, 1, 1, 1, 1}))
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int
	work := make([]int, len(nums))
	var p func(int, int)
	mask := 1<<uint(len(nums)) - 1
	p = func(i int, used int) {
		if i == len(nums) {
			res = append(res, append([]int{}, work...))
			return
		}
		rem := mask &^ used
		last, first := 0, true
		for rem != 0 {
			j := bits.TrailingZeros(uint(rem))
			bit := 1 << uint(j)
			rem = rem &^ bit
			if !first && last == nums[j] {
				continue
			}
			last, first = nums[j], false
			work[i] = nums[j]
			p(i+1, used|bit)
		}
	}
	p(0, 0)
	return res
}
