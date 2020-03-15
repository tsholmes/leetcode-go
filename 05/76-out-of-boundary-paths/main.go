package main

func main() {

}

func findPaths(m int, n int, N int, i int, j int) int {
	res := 0
	cur := map[[2]int]int{[2]int{i, j}: 1}
	dirs := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	const mod = 1e9 + 7
	for move := 0; move < N; move++ {
		next := map[[2]int]int{}
		for k, v := range cur {
			for _, d := range dirs {
				i2, j2 := k[0]+d[0], k[1]+d[1]
				k2 := [2]int{i2, j2}
				if i2 < 0 || i2 >= m || j2 < 0 || j2 >= n {
					res = (res + v) % mod
				} else {
					next[k2] = (next[k2] + v) % mod
				}
			}
		}
		cur = next
	}
	return res
}
