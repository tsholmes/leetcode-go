package main

import "fmt"

func main() {
	fmt.Println(ways([]string{".A", "AA", "A."}, 3))
}

func ways(pizza []string, k int) int {
	H, W := len(pizza), len(pizza[0])

	counts := make([][]int, H)
	for i := range counts {
		counts[i] = make([]int, W)
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if pizza[i][j] == 'A' {
				counts[i][j]++
			}
			if i == 0 && j > 0 {
				counts[i][j] += counts[i][j-1]
			} else if j == 0 && i > 0 {
				counts[i][j] += counts[i-1][j]
			} else if i > 0 && j > 0 {
				counts[i][j] += counts[i-1][j] + counts[i][j-1] - counts[i-1][j-1]
			}
		}
	}

	getc := func(i1, j1, i2, j2 int) int {
		f := counts[i2][j2]
		l, t, tl := 0, 0, 0
		if i1 > 0 && j1 > 0 {
			l = counts[i1-1][j2]
			t = counts[i2][j1-1]
			tl = counts[i1-1][j1-1]
		} else if i1 > 0 {
			l = counts[i1-1][j2]
		} else if j1 > 0 {
			t = counts[i2][j1-1]
		}
		return f - l - t + tl
	}

	const mod = 1e9 + 7

	memo := map[[3]int]int{}
	var search func(i, j, krem int) int
	search = func(i, j, krem int) int {
		if krem == 1 {
			if getc(i, j, H-1, W-1) > 0 {
				return 1
			}
			return 0
		}
		key := [3]int{i, j, krem}
		if v, ok := memo[key]; ok {
			return v
		}

		res := 0

		for i2 := i + 1; i2 < H; i2++ {
			if getc(i, j, i2-1, W-1) > 0 {
				res = (res + search(i2, j, krem-1)) % mod
			}
		}

		for j2 := j + 1; j2 < W; j2++ {
			if getc(i, j, H-1, j2-1) > 0 {
				res = (res + search(i, j2, krem-1)) % mod
			}
		}

		memo[key] = res
		return res
	}

	r := search(0, 0, k)
	return r
}
