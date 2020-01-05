package util

// This is meant to be copied into a solution

type state [3]int
type key [2]int

func bfs() (state, bool) {
	initState := state{0, 0, 0}
	valid := func(s state) bool {
		return s[0]+s[1] < 5 && s[0]-s[1] < 5 && -s[0]+s[1] < 5 && -s[0]-s[1] < 5
	}
	keyOf := func(s state) key {
		return key{s[0], s[1]}
	}
	isEnd := func(s state) bool {
		return false
	}
	onState := func(s state) {
		// do something for each state
	}
	next := func(s state, f func(state)) {
		f(state{s[0] + 1, s[1], s[2] + 1})
		f(state{s[0], s[1] + 1, s[2] + 1})
		f(state{s[0] - 1, s[1], s[2] + 1})
		f(state{s[0], s[1] - 1, s[2] + 1})
	}

	queue := []state{initState}
	seen := map[key]bool{}
	for len(queue) > 0 {
		// for dfs change to
		// s := queue[len(queue)-1]
		// queue = queue[:len(queue)-1]
		s := queue[0]
		queue = queue[1:]
		k := keyOf(s)
		if seen[k] {
			continue
		}
		seen[k] = true

		onState(s)
		if isEnd(s) {
			return s, true
		}

		next(s, func(n state) {
			if !valid(n) {
				return
			}
			if seen[keyOf(n)] {
				return
			}
			queue = append(queue, n)
		})
	}

	return state{}, false
}
