package main

import "fmt"

func main() {
	fmt.Println(numberOfBoomerangs([][]int{{0, 0}, {1, 0}, {2, 0}}))
}

func numberOfBoomerangs(points [][]int) int {
	res := 0
	for _, p := range points {
		ds := map[int]int{}
		for _, p2 := range points {
			dx, dy := p2[0]-p[0], p2[1]-p[1]
			ds[dx*dx+dy*dy]++
		}
		for _, v := range ds {
			res += v * (v - 1)
		}
	}
	return res
}
