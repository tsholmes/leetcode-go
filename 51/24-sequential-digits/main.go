package main

import "fmt"

func main() {
	fmt.Println(sequentialDigits(1000, 13000))
}

func toDigits(val int) []int {
	res := []int{}
	for val > 0 {
		res = append([]int{val % 10}, res...)
		val = val / 10
	}
	return res
}

func sequentialDigits(low int, high int) []int {
	lowd := toDigits(low)
	highd := toDigits(high)
	if len(lowd) < len(highd) {
		l := make([]int, len(highd))
		copy(l[len(highd)-len(lowd):], lowd)
		lowd = l
	}
	var res []int
	var search func(int)
	work := make([]int, len(lowd))
	search = func(i int) {
		if i == len(lowd) {
			val := 0
			for _, d := range work {
				val *= 10
				val += d
			}
			if val >= low && val <= high {
				res = append(res, val)
			}
			return
		}
		if i == 0 || work[i-1] == 0 {
			for d := lowd[i]; d <= 9; d++ {
				work[i] = d
				search(i + 1)
			}
		} else {
			if work[i-1] == 9 {
				return
			}
			work[i] = work[i-1] + 1
			search(i + 1)
		}
	}
	search(0)
	return res
}
