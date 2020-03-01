package main

import "sort"

func main() {

}

func rankTeams(votes []string) string {
	N := len(votes[0])

	votecs := map[byte][]int{}
	for _, v := range votes {
		for i := 0; i < N; i++ {
			c := v[i]
			if _, ok := votecs[c]; !ok {
				votecs[c] = make([]int, N)
			}
			votecs[c][i]++
		}
	}

	var all []byte
	for k := range votecs {
		all = append(all, k)
	}
	sort.Slice(all, func(i, j int) bool {
		vi, vj := votecs[all[i]], votecs[all[j]]
		for k := 0; k < N; k++ {
			if vi[k] > vj[k] {
				return true
			}
			if vi[k] < vj[k] {
				return false
			}
		}
		return all[i] < all[j]
	})
	return string(all)
}
