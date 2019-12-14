package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{1, 2, 3, 4}))
}

func permute(nums []int) [][]int {
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
		for rem != 0 {
			j := bits.TrailingZeros(uint(rem))
			bit := 1 << uint(j)
			work[i] = nums[j]
			p(i+1, used|bit)
			rem = rem &^ bit
		}
	}
	p(0, 0)
	return res
}
