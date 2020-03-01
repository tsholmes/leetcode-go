package main

import "fmt"

func main() {
	fmt.Println(makesquare([]int{1, 1, 2, 2, 2}))
}

func makesquare(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%4 != 0 {
		return false
	}
	if sum == 0 {
		return false
	}
	side := sum / 4

	valid := func(ns []int) bool {
		for mask := 0; mask < (1 << uint(len(ns))); mask++ {
			sa, sb := 0, 0
			for i, n := range ns {
				if mask&(1<<uint(i)) != 0 {
					sa += n
				} else {
					sb += n
				}
			}
			if sa == side && sb == side {
				return true
			}
		}
		return false
	}

	var na []int
	var nb []int
	for mask := 0; mask < (1 << uint(len(nums))); mask++ {
		na, nb = na[:0], nb[:0]
		sa, sb := 0, 0
		for i, n := range nums {
			if mask&(1<<uint(i)) != 0 {
				na = append(na, n)
				sa += n
			} else {
				nb = append(nb, n)
				sb += n
			}
		}
		if sa != side<<1 || sb != side<<1 {
			continue
		}
		if valid(na) && valid(nb) {
			return true
		}
	}

	return false
}
