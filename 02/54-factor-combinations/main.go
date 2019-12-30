package main

import "fmt"

func main() {
	fmt.Println(getFactors(1))
	fmt.Println(getFactors(37))
	fmt.Println(getFactors(12))
	fmt.Println(getFactors(32))
}

func getFactors(n int) [][]int {
	var res [][]int

	var work []int
	var search func(int, int)
	search = func(cur int, last int) {
		if cur == 1 {
			if len(work) > 0 {
				res = append(res, append([]int{}, work...))
			}
			return
		}
		wlen := len(work)
		for f := last; f*f <= cur; f++ {
			if cur%f != 0 {
				continue
			}
			work = append(work, f)
			search(cur/f, f)
			work = work[:wlen]
		}
		if cur != n {
			work = append(work, cur)
			search(1, cur)
			work = work[:wlen]
		}
	}

	search(n, 2)
	return res
}
