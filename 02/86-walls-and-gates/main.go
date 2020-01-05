package main

func main() {

}

func wallsAndGates(rooms [][]int) {
	fill := func(gpos [2]int) {
		type state [3]int
		type key [2]int
		initState := state{gpos[0], gpos[1], 0}
		valid := func(s state) bool {
			i, j := s[0], s[1]
			return i >= 0 && i < len(rooms) && j >= 0 && j < len(rooms[i]) && rooms[i][j] > 0
		}
		keyOf := func(s state) key {
			return key{s[0], s[1]}
		}
		onState := func(s state) {
			i, j := s[0], s[1]
			if rooms[i][j] > s[2] {
				rooms[i][j] = s[2]
			}
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
			s := queue[0]
			queue = queue[1:]
			k := keyOf(s)
			if seen[k] {
				continue
			}
			seen[k] = true

			onState(s)

			next(s, func(n state) {
				if !valid(n) {
					return
				}
				queue = append(queue, n)
			})
		}
	}

	for i := range rooms {
		for j := range rooms[i] {
			if rooms[i][j] == 0 {
				fill([2]int{i, j})
			}
		}
	}
}
