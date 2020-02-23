package main

func main() {

}

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if desiredTotal == 0 {
		return true
	}
	if (maxChoosableInteger*(maxChoosableInteger+1))/2 < desiredTotal {
		return false
	}

	type state struct {
		mask int
		turn bool
	}
	memo := map[state]bool{}
	var search func(mask int, sum int, turn bool) bool
	search = func(mask int, sum int, turn bool) bool {
		if sum >= desiredTotal {
			return turn
		}
		k := state{mask: mask, turn: turn}
		if v, ok := memo[k]; ok {
			return v
		}

		var res bool

		if turn {
			all := true
			for i := 1; i <= maxChoosableInteger && all; i++ {
				bit := 1 << uint(i-1)
				if mask&bit != 0 {
					continue
				}
				all = search(mask|bit, sum+i, !turn)
			}
			res = all
		} else {
			any := false
			for i := 1; i <= maxChoosableInteger && !any; i++ {
				bit := 1 << uint(i-1)
				if mask&bit != 0 {
					continue
				}
				any = search(mask|bit, sum+i, !turn)
			}
			res = any
		}
		memo[k] = res
		return res
	}
	return search(0, 0, false)
}
