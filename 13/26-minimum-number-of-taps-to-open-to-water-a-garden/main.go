package main

import "fmt"

func main() {
	fmt.Println(minTaps(5, []int{3, 4, 1, 1, 0, 0}))
	fmt.Println(minTaps(3, []int{0, 0, 0, 0}))
	fmt.Println(minTaps(7, []int{1, 2, 1, 0, 2, 1, 0, 1}))
	fmt.Println(minTaps(8, []int{4, 0, 0, 0, 4, 0, 0, 0, 4}))
}

func minTaps(n int, ranges []int) int {
	cover := make([][]int, n)
	for i := range cover {
		for j, l := range ranges {
			if i >= j-l && (i+1) <= j+l {
				cover[i] = append(cover[i], j)
			}
		}
	}
	memo := map[int]int{}
	var search func(int) int
	search = func(p int) int {
		if p >= n {
			return 0
		}
		if v, ok := memo[p]; ok {
			return v
		}
		max := p
		for _, j := range cover[p] {
			if j+ranges[j] > max {
				max = j + ranges[j]
			}
		}
		if max <= p {
			return -1
		}
		r := search(max)
		if r >= 0 {
			r++
		}
		memo[p] = r
		return r
	}

	return search(0)
}
