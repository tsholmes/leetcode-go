package main

func main() {

}

func numberOfPatterns(m int, n int) int {
	edges := map[int][]int{
		1: {2, 4, 6, 5, 8},
		2: {1, 3, 4, 5, 6, 7, 9},
		3: {2, 4, 5, 6, 8},
		4: {1, 2, 3, 5, 7, 8, 9},
		5: {1, 2, 3, 4, 6, 7, 8, 9},
		6: {1, 2, 3, 5, 7, 8, 9},
		7: {2, 4, 5, 6, 8},
		8: {1, 3, 4, 5, 6, 7, 9},
		9: {2, 4, 5, 6, 8},
	}
	skipEdges := map[int][][2]int{
		1: {{3, 2}, {9, 5}, {7, 4}},
		2: {{8, 5}},
		3: {{1, 2}, {7, 5}, {9, 6}},
		4: {{6, 5}},
		5: {},
		6: {{4, 5}},
		7: {{1, 4}, {3, 5}, {9, 8}},
		8: {{2, 5}},
		9: {{7, 8}, {1, 5}, {3, 6}},
	}

	memo := map[[2]int]int{}
	var search func(int, int, int) int
	search = func(cur int, used int, count int) int {
		if count > n {
			return 0
		}
		k := [2]int{cur, used}
		if v, ok := memo[k]; ok {
			return v
		}

		var res int
		if count >= m {
			res++
		}
		for _, e := range edges[cur] {
			bit := 1 << uint(e)
			if used&bit == 0 {
				res += search(e, used|bit, count+1)
			}
		}
		for _, ee := range skipEdges[cur] {
			b1, b2 := 1<<uint(ee[0]), 1<<uint(ee[1])
			if used&b2 != 0 && used&b1 == 0 {
				res += search(ee[0], used|b1, count+1)
			}
		}
		memo[k] = res
		return res
	}

	return search(1, 1<<1, 1)*4 + search(2, 1<<2, 1)*4 + search(5, 1<<5, 1)
}
