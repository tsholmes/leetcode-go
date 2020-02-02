package main

func main() {
}

func sequenceReconstruction(org []int, seqs [][]int) bool {
	edges := map[int][]int{}
	for _, seq := range seqs {
		for _, n := range seq {
			if _, ok := edges[n]; !ok {
				edges[n] = nil
			}
		}
		for i := 1; i < len(seq); i++ {
			edges[seq[i-1]] = append(edges[seq[i-1]], seq[i])
		}
	}

	depths := map[int]int{}
	var search func(n int, d int)
	used := map[int]bool{}
	valid := true
	search = func(n int, d int) {
		if !valid {
			return
		}
		if v, ok := depths[n]; ok && v >= d {
			return
		}
		if used[n] {
			valid = false
			return
		}
		used[n] = true
		depths[n] = d
		for _, n2 := range edges[n] {
			search(n2, d+1)
		}
		used[n] = false
	}

	for k := range edges {
		search(k, 0)
	}

	if !valid {
		return false
	}

	if len(depths) != len(org) {
		return false
	}

	for i, n := range org {
		if v, ok := depths[n]; !ok || v != i {
			return false
		}
	}

	return true
}
