package main

import "fmt"

func main() {
	fmt.Println(numOfArrays(2, 3, 1))
	fmt.Println(numOfArrays(9, 1, 1))
}

func numOfArrays(n int, m int, k int) int {
	const mod = 1e9 + 7
	memo := map[[3]int]int{}

	var npow [101][101]int
	npow[0][0] = 1
	for i := 1; i <= 100; i++ {
		npow[0][i] = 1
		v := 1
		npow[i][0] = v
		for j := 1; j <= 100; j++ {
			v = (v * i) % mod
			npow[i][j] = v
		}
	}

	var search func(lastv, rem, krem int) int
	search = func(lastv, rem, krem int) int {
		if rem < krem {
			return 0
		}
		if krem == 0 {
			return npow[lastv][rem]
		}
		key := [3]int{lastv, rem, krem}
		if v, ok := memo[key]; ok {
			return v
		}

		var res int
		for next := lastv + 1; next <= m; next++ {
			for fill := 0; fill < rem; fill++ {
				if lastv == 0 && fill > 0 {
					continue
				}
				res = (res + npow[lastv][fill]*search(next, rem-fill-1, krem-1)) % mod
			}
		}
		memo[key] = res

		return res
	}

	r := search(0, n, k)
	return r
}
