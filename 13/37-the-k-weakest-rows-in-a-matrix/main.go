package main

import "sort"

func main() {

}
func kWeakestRows(mat [][]int, k int) []int {
	ns := make([]int, len(mat))
	for i := range ns {
		ns[i] = i
	}
	sort.SliceStable(ns, func(i, j int) bool {
		si, sj := 0, 0
		for _, n := range mat[ns[i]] {
			si += n
		}
		for _, n := range mat[ns[j]] {
			sj += n
		}
		return si < sj
	})
	return ns[:k]
}
