package main

import "sort"

func main() {

}

func minSetSize(arr []int) int {
	counts := map[int]int{}
	for _, n := range arr {
		counts[n]++
	}

	var nns []int
	for k := range counts {
		nns = append(nns, k)
	}
	sort.Slice(nns, func(i, j int) bool {
		return counts[nns[i]] > counts[nns[j]]
	})

	count := 0
	rem := 0
	for count*2 < len(arr) {
		count += counts[nns[0]]
		rem++
		nns = nns[1:]
	}
	return rem
}
