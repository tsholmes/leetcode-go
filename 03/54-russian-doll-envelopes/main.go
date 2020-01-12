package main

import "sort"

func main() {

}

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		return (envelopes[i][0] * envelopes[i][1]) <= (envelopes[j][0] * envelopes[j][1])
	})
	edges := make([][]int, len(envelopes))
	for i, ei := range envelopes {
		for j, ej := range envelopes {
			if ei[0] < ej[0] && ei[1] < ej[1] {
				edges[i] = append(edges[i], j)
			}
		}
	}

	max := 0
	depths := make([]int, len(envelopes))
	var search func(int, int)
	search = func(i int, depth int) {
		if depths[i] >= depth {
			return
		}
		depths[i] = depth
		if depth > max {
			max = depth
		}
		for _, j := range edges[i] {
			search(j, depth+1)
		}
	}
	for i := range edges {
		search(i, 1)
	}
	return max
}
