package main

import "sort"

func main() {

}

func longestDiverseString(a int, b int, c int) string {
	counts := map[byte]int{'a': a, 'b': b, 'c': c}
	order := []byte{'a', 'b', 'c'}
	cur := map[byte]int{}
	var last byte
	var res []byte
	for iter := 0; iter < a+b+c; iter++ {
		sort.Slice(order, func(i, j int) bool { return counts[order[i]] > counts[order[j]] })
		var nc byte
		for _, b := range order {
			if counts[b] > 0 && cur[b] < 2 {
				nc = b
				break
			}
		}
		if nc == 0 {
			break
		}
		res = append(res, nc)
		counts[nc]--
		if nc != last {
			cur = map[byte]int{}
		}
		cur[nc]++
		last = nc
	}
	return string(res)
}
