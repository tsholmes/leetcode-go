package main

func main() {

}

func numTeams(rating []int) int {
	type state struct {
		i, r int
		gt   bool
	}
	memo := map[state]int{}

	var search func(i int, r int, gt bool) int
	search = func(i, r int, gt bool) int {
		if r == 0 {
			return 1
		}
		k := state{i, r, gt}
		if v, ok := memo[k]; ok {
			return v
		}

		res := 0
		for j := i + 1; j < len(rating); j++ {
			if (rating[i] > rating[j] && gt) || (rating[i] < rating[j] && !gt) {
				res += search(j, r-1, gt)
			}
		}

		memo[k] = res
		return res
	}

	tres := 0
	for i := range rating {
		tres += search(i, 2, true) + search(i, 2, false)
	}
	return tres
}
