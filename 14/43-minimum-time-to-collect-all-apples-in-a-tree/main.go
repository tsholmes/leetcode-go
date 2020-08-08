package main

func main() {

}

func minTime(n int, edges [][]int, hasApple []bool) int {
	es := map[int][]int{}
	for _, e := range edges {
		es[e[0]] = append(es[e[0]], e[1])
		es[e[1]] = append(es[e[1]], e[0])
	}

	var walk func(n, p int) (int, bool)
	walk = func(n, p int) (int, bool) {
		has := hasApple[n]
		res := 0
		for _, n2 := range es[n] {
			if n2 == p {
				continue
			}
			a, h := walk(n2, n)
			if h {
				res += 2 + a
				has = true
			}
		}
		return res, has
	}

	r, _ := walk(0, -1)
	return r
}
