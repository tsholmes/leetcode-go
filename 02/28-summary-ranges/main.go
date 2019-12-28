package main

import (
	"fmt"
	"strconv"
)

func main() {

}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	var res []string
	min, max := nums[0], nums[0]
	for _, n := range nums[1:] {
		if n == max+1 {
			max = n
		} else {
			if min == max {
				res = append(res, strconv.Itoa(min))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", min, max))
			}
			min, max = n, n
		}
	}
	if min == max {
		res = append(res, strconv.Itoa(min))
	} else {
		res = append(res, fmt.Sprintf("%d->%d", min, max))
	}
	return res
}
