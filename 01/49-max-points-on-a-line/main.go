package main

import "fmt"

func main() {
	fmt.Println(maxPoints([][]int{{1, 1}, {2, 2}, {3, 3}}))
	fmt.Println(maxPoints([][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}))
	fmt.Println(maxPoints([][]int{{-3, -5}, {-1, -3}, {1, -1}, {3, 1}}))
	fmt.Println(maxPoints([][]int{{1, 1}, {1, 1}, {2, 3}}))
}

func gcf(a, b int) int {
	if b == 0 {
		return a
	}
	return gcf(b, a%b)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func maxPoints(points [][]int) int {
	if len(points) <= 2 {
		return len(points)
	}
	xs := map[int]int{}
	counts := map[[4]int]int{}
	for i, p1 := range points {
		xs[p1[0]]++
		seen := map[[4]int]bool{}
		isNew := map[[4]int]bool{}
		for _, p2 := range points[:i] {
			if p1[0] == p2[0] {
				continue
			}
			dx, dy := p2[0]-p1[0], p2[1]-p1[1]
			f := gcf(dx, dy)
			dx /= f
			dy /= f
			m := p1[0] % dx
			if m < 0 {
				m += dx
			}
			yi := p1[1] - (p1[0]/dx)*dy
			key := [4]int{dx, dy, m, yi}
			_, exists := counts[key]
			if !exists {
				counts[key] = 2
				isNew[key] = true
			} else if isNew[key] || !seen[key] {
				counts[key]++
			}
			seen[key] = true
		}
	}
	max := 0
	for _, v := range xs {
		if v > max {
			max = v
		}
	}
	for _, v := range counts {
		if v > max {
			max = v
		}
	}
	return max
}
