package main

func main() {

}

func integerReplacement(n int) int {
	memo := map[int]int{}
	var search func(int) int
	search = func(i int) int {
		if i == 1 {
			return 0
		}
		if v, ok := memo[i]; ok {
			return v
		}
		var res int
		if i&1 == 0 {
			res = 1 + search(i/2)
		} else {
			a, b := search(i-1), search(i+1)
			if a < b {
				res = a + 1
			} else {
				res = b + 1
			}
		}
		memo[i] = res
		return res
	}
	return search(n)
}
