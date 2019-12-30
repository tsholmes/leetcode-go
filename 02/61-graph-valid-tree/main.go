package main

func main() {

}

func validTree(n int, edges [][]int) bool {
	ps := make([]int, n)
	for i := range ps {
		ps[i] = i
	}
	var (
		get  func(int) int
		join func(int, int)
	)
	get = func(i int) int {
		if ps[i] == i {
			return i
		}
		p := get(ps[i])
		ps[i] = p
		return p
	}
	join = func(i, j int) {
		pi, pj := get(i), get(j)
		ps[pi] = pj
	}
	for _, e := range edges {
		if get(e[0]) == get(e[1]) {
			return false
		}
		join(e[0], e[1])
	}
	for i := range ps {
		if get(i) != get(0) {
			return false
		}
	}
	return true
}
