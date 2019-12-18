package main

import "fmt"

func main() {
	fmt.Println(findMissingRanges([]int{0, 1, 3, 50, 75}, 0, 99))
}

func findMissingRanges(nums []int, lower int, upper int) []string {
	var res []string
	last := lower - 1
	for _, n := range nums {
		if n > last+1 {
			low := last + 1
			high := n - 1
			if low == high {
				res = append(res, fmt.Sprintf("%d", low))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", low, high))
			}
		}
		last = n
	}
	if last < upper {
		low := last + 1
		high := upper
		if low == high {
			res = append(res, fmt.Sprintf("%d", low))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", low, high))
		}
	}
	return res
}
