package main

func main() {

}

func checkRecord(n int) int {
	const mod = 1e9 + 7
	memo := map[[3]int]int{}
	var search func(i, ac, lc int) int
	search = func(i, ac, lc int) int {
		if ac > 1 {
			return 0
		}
		if lc > 2 {
			return 0
		}
		if i == n {
			return 1
		}
		k := [3]int{i, ac, lc}
		if v, ok := memo[k]; ok {
			return v
		}

		res := (search(i+1, ac, 0) + search(i+1, ac+1, 0) + search(i+1, ac, lc+1)) % mod

		memo[k] = res
		return res
	}
	return search(0, 0, 0)
}
