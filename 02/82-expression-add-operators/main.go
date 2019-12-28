package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	do := func(num string, target int) {
		start := time.Now()
		addOperators(num, target)
		end := time.Now()
		fmt.Println(end.Sub(start))
	}
	do("123", 6)
	do("232", 8)
	do("105", 5)
	do("00", 0)
	do("3456237490", 9191)
	do("22222222222", 0)
}

func addOperators(num string, target int) []string {
	var res []string

	var search func(int, int, int)
	var work []byte
	search = func(base int, last int, pos int) {
		if pos == len(num) {
			if base+last == target {
				res = append(res, string(work))
			}
			return
		}
		wlen := len(work)
		for i := pos + 1; i <= len(num); i++ {
			if num[pos] == '0' && i-pos > 1 {
				continue
			}
			v, _ := strconv.Atoi(num[pos:i])

			// add
			work = work[:wlen]
			if pos > 0 {
				work = append(work, '+')
			}
			work = append(work, num[pos:i]...)
			search(base+last, v, i)

			if pos > 0 {
				// sub
				work = work[:wlen]
				work = append(work, '-')
				work = append(work, num[pos:i]...)
				search(base+last, -v, i)

				// mul
				work = work[:wlen]
				work = append(work, '*')
				work = append(work, num[pos:i]...)
				search(base, last*v, i)
			}
		}
	}

	search(0, 0, 0)

	return res
}
