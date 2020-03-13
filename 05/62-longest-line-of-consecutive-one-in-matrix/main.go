package main

func main() {

}

func longestLine(M [][]int) int {
	if len(M) == 0 || len(M[0]) == 0 {
		return 0
	}
	A, B := len(M), len(M[0])
	dirs := [][2]int{{1, 0}, {0, 1}, {1, 1}, {1, -1}}
	edges := [][3]bool{{true, false, false}, {false, true, false}, {true, true, false}, {true, false, true}}

	max := 0
	run := func(i, j int, d [2]int) {
		cur := 0
		for i >= 0 && i < A && j >= 0 && j < B {
			if M[i][j] == 1 {
				cur++
				if cur > max {
					max = cur
				}
			} else {
				cur = 0
			}
			i, j = i+d[0], j+d[1]
		}
	}

	for di := 0; di < 4; di++ {
		d, e := dirs[di], edges[di]
		if e[0] {
			for j := 0; j < B; j++ {
				run(0, j, d)
			}
		}
		if e[1] {
			for i := 0; i < A; i++ {
				run(i, 0, d)
			}
		}
		if e[2] {
			for i := 0; i < A; i++ {
				run(i, B-1, d)
			}
		}
	}

	return max
}
