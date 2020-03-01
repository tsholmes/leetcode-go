package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumDistance("CAKE"))
	fmt.Println(minimumDistance("HAPPY"))
	fmt.Println(minimumDistance("NEW"))
	fmt.Println(minimumDistance("YEAR"))
	fmt.Println(minimumDistance("GSALRVPVEYXICXWTWIZTXHAEGXSVANEEPRBQKHVFOKVOHUUXIVWGDLPILRSJGJRMCFMDSUDQEWOENIXRGYZFXCIOKMQBLMXXMREYWWRSEUBJTOQYSNGOGMCDRMDQVOXUFBGHHDTMNCRGQGGPCGDQETGRSQUTNTMUIVZPUUTZALLYZGWUBGOGQVGAEHOPCTIUJMKXXVWJVCTDPLTTUCHENHAPJEQULXVEDQKKVXCLZGGJQWNUXXMMBEOCJUVEQGTPLESES"))
}

// submitted version
func minimumDistance2(word string) int {
	board := [][]byte{
		[]byte(`ABCDEF`),
		[]byte(`GHIJKL`),
		[]byte(`MNOPQR`),
		[]byte(`STUVWX`),
		[]byte(`YZ`),
	}
	cps := map[byte][2]int{}
	for i, r := range board {
		for j, c := range r {
			cps[c] = [2]int{i, j}
		}
	}

	dist := func(i, j int) int {
		if i == 0 {
			return 0
		}
		p1 := cps[word[i-1]]
		p2 := cps[word[j-1]]
		dx := p1[0] - p2[0]
		dy := p1[1] - p2[1]
		if dx < 0 {
			dx = -dx
		}
		if dy < 0 {
			dy = -dy
		}
		return dx + dy
	}

	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}

	best := math.MaxInt32

	dp := make([][][]int, len(word))
	for i := range dp {
		dp[i] = make([][]int, len(word)+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, len(word)+1)
			for k := range dp[i][j] {
				dp[i][j][k] = math.MaxInt32
			}
		}
	}
	for i := range dp {
		dp[0][i+1][0] = 0
		dp[0][0][i+1] = 0
	}
	for i := 1; i < len(word); i++ {
		for j := 0; j <= i; j++ {
			mk := 0
			if j != i {
				mk = i
			}
			for k := mk; k <= i; k++ {
				dp[i][i+1][k] = min(dp[i][i+1][k], dp[i-1][j][k]+dist(j, i+1))
				dp[i][j][i+1] = min(dp[i][j][i+1], dp[i-1][j][k]+dist(k, i+1))
				if i == len(word)-1 {
					best = min(best, dp[i][i+1][k])
					best = min(best, dp[i][j][i+1])
				}
			}
		}
	}

	return best
}

// shorter memoized search version
func minimumDistance(word string) int {
	board := [][]byte{
		[]byte(`ABCDEF`),
		[]byte(`GHIJKL`),
		[]byte(`MNOPQR`),
		[]byte(`STUVWX`),
		[]byte(`YZ`),
	}
	cps := map[byte][2]int{}
	for i, r := range board {
		for j, c := range r {
			cps[c] = [2]int{i, j}
		}
	}

	dist := func(i, j int) int {
		if i == 0 {
			return 0
		}
		p1 := cps[word[i-1]]
		p2 := cps[word[j-1]]
		dx := p1[0] - p2[0]
		dy := p1[1] - p2[1]
		if dx < 0 {
			dx = -dx
		}
		if dy < 0 {
			dy = -dy
		}
		return dx + dy
	}

	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}

	memo := map[[3]int]int{}
	var search func(i, l, r int) int
	search = func(i, l, r int) int {
		if i == len(word) {
			return 0
		}
		k := [3]int{i, l, r}
		if v, ok := memo[k]; ok {
			return v
		}

		memo[k] = min(search(i+1, i+1, r)+dist(l, i+1), search(i+1, l, i+1)+dist(r, i+1))
		return memo[k]
	}

	return search(0, 0, 0)
}
