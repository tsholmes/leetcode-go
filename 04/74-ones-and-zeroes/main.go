package main

func main() {

}

func findMaxForm(strs []string, m int, n int) int {
	var counts [][2]int
	for _, s := range strs {
		var cs [2]int
		for i := 0; i < len(s); i++ {
			if s[i] == '1' {
				cs[1]++
			} else {
				cs[0]++
			}
		}
		counts = append(counts, cs)
	}

	memo := map[[3]int]int{}
	var search func(zs, os, i int) int
	search = func(zs, os, i int) int {
		if i == len(strs) {
			return 0
		}
		k := [3]int{zs, os, i}
		if v, ok := memo[k]; ok {
			return v
		}
		max := search(zs, os, i+1)
		if counts[i][0] <= zs && counts[i][1] <= os {
			m2 := 1 + search(zs-counts[i][0], os-counts[i][1], i+1)
			if m2 > max {
				max = m2
			}
		}
		memo[k] = max
		return max
	}

	return search(m, n, 0)
}
