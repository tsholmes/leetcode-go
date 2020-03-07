package main

func main() {

}

func PredictTheWinner(nums []int) bool {
	type state struct {
		s1, s2 int
		l, r   int
		turn   bool
	}
	memo := map[state]bool{}
	var search func(s1, s2, l, r int, turn bool) bool
	search = func(s1, s2, l, r int, turn bool) bool {
		if l == r {
			return s1 >= s2
		}
		k := state{s1, s2, l, r, turn}
		if v, ok := memo[k]; ok {
			return v
		}

		var res bool
		if turn {
			res = search(s1+nums[l], s2, l+1, r, false) || search(s1+nums[r-1], s2, l, r-1, false)
		} else {
			res = search(s1, s2+nums[l], l+1, r, true) && search(s1, s2+nums[r-1], l, r-1, true)
		}
		memo[k] = res
		return res
	}
	return search(0, 0, 0, len(nums), true)
}
