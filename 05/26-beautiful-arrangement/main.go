package main

func main() {

}

func countArrangement(N int) int {
	memo := map[[2]int]int{}
	var search func(i int, mask int) int
	search = func(i int, mask int) int {
		if i == N {
			return 1
		}
		k := [2]int{i, mask}
		if v, ok := memo[k]; ok {
			return v
		}
		res := 0
		for j := 0; j < N; j++ {
			bit := 1 << uint(j)
			if mask&bit != 0 {
				continue
			}
			if (i+1)%(j+1) != 0 && (j+1)%(i+1) != 0 {
				continue
			}
			res += search(i+1, mask|bit)
		}
		memo[k] = res
		return res
	}
	return search(0, 0)
}
