package util

import "container/heap"

// This is meant to be copied into a solution

func bfs() {
	type state [3]int
	type key [2]int
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
	var res state
	var found bool
	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]
		k := keyOf(s)
		if seen[k] {
			continue
		}
		seen[k] = true

		onState(s)
		if isEnd(s) {
			res = s
			found = true
			break
		}

		next(s, func(n state) {
			if !valid(n) {
				return
			}
			queue = append(queue, n)
		})
	}

	_, _ = res, found
}

func dfs() {
	type state [3]int
	type key [2]int
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
	var res state
	var found bool
	for len(queue) > 0 {
		s := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		k := keyOf(s)
		if seen[k] {
			continue
		}
		seen[k] = true

		onState(s)
		if isEnd(s) {
			res = s
			found = true
			break
		}

		next(s, func(n state) {
			if !valid(n) {
				return
			}
			queue = append(queue, n)
		})
	}

	_, _ = res, found
}

type dstate [2]int
type stateHeap []dstate

func (h *stateHeap) Len() int           { return len(*h) }
func (h *stateHeap) Less(i, j int) bool { return (*h)[i][1] < (*h)[j][1] }
func (h *stateHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *stateHeap) Push(x interface{}) { *h = append(*h, x.(dstate)) }
func (h *stateHeap) Pop() interface{}   { v := (*h)[h.Len()-1]; *h = (*h)[:h.Len()-1]; return v }

func djikstras() {
	type key int
	initState := dstate{0, 0}
	edges := [][]bool{
		{false, true, true, true, true},
		{true, false, true, true, true},
		{true, true, false, true, true},
		{true, true, true, false, true},
		{true, true, true, true, false},
	}
	valid := func(s dstate) bool {
		return true
	}
	keyOf := func(s dstate) key {
		return key(s[0])
	}
	isEnd := func(s dstate) bool {
		return false
	}
	onState := func(s dstate) {
		// do something for each state
	}
	next := func(s dstate, f func(dstate)) {
		cost := func(i, j int) int {
			return (i - j) * (i - j)
		}
		for i, ok := range edges[s[0]] {
			if ok {
				f(dstate{i, s[1] + cost(s[0], i)})
			}
		}
	}

	queue := stateHeap{initState}
	seen := map[key]bool{}
	var res dstate
	var found bool
	for len(queue) > 0 {
		s := heap.Pop(&queue).(dstate)
		k := keyOf(s)
		if seen[k] {
			continue
		}
		seen[k] = true

		onState(s)
		if isEnd(s) {
			res = s
			found = true
			break
		}

		next(s, func(n dstate) {
			if !valid(n) {
				return
			}
			heap.Push(&queue, n)
		})
	}

	_, _ = res, found
}
