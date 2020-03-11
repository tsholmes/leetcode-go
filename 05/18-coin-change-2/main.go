package main

func main() {

}

func change(amount int, coins []int) int {
	memo := map[[2]int]int{}
	var search func(int, int) int
	search = func(n int, mini int) int {
		if n == 0 {
			return 1
		}
		if n < 0 {
			return 0
		}
		k := [2]int{n, mini}
		if v, ok := memo[k]; ok {
			return v
		}
		count := 0
		for i := mini; i < len(coins); i++ {
			count += search(n-coins[i], i)
		}
		memo[k] = count
		return count
	}
	return search(amount, 0)
}
