package main

import "fmt"

func main() {
	fmt.Println(frogPosition(7, [][]int{
		{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 6},
	}, 2, 4))
}

func frogPosition(n int, edges [][]int, t int, target int) float64 {
	es := map[int][]int{}
	for _, e := range edges {
		es[e[0]] = append(es[e[0]], e[1])
		es[e[1]] = append(es[e[1]], e[0])
	}

	var prob float64
	visited := map[int]bool{}
	var search func(i int, p float64, tt int)
	search = func(i int, p float64, tt int) {
		if tt == t {
			if i == target {
				prob += p
			}
			return
		}
		visited[i] = true
		ec := 0
		for _, e := range es[i] {
			if !visited[e] {
				ec++
			}
		}
		for _, e := range es[i] {
			if !visited[e] {
				search(e, p/float64(ec), tt+1)
			}
		}
		if ec == 0 && i == target {
			prob += p
		}
		visited[i] = false
	}
	search(1, 1.0, 0)
	return prob
}
