package main

import "sort"

func main() {

}

func findItinerary(tickets [][]string) []string {
	type tick struct {
		next string
		idx  int
	}
	es := map[string][]tick{}
	for i, t := range tickets {
		es[t[0]] = append(es[t[0]], tick{t[1], i})
	}
	for k := range es {
		sort.Slice(es[k], func(i, j int) bool {
			return es[k][i].next < es[k][j].next
		})
	}

	var found bool
	var res []string
	used := make([]bool, len(tickets))
	var search func(string)
	search = func(cur string) {
		res = append(res, cur)
		if len(res) == len(tickets)+1 {
			found = true
			return
		}
		rlen := len(res)
		for _, t := range es[cur] {
			if used[t.idx] {
				continue
			}
			used[t.idx] = true
			search(t.next)
			if found {
				return
			}
			used[t.idx] = false
			res = res[:rlen]
		}
	}
	search("JFK")
	return res
}
