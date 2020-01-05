package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(watchedVideosByFriends([][]string{
		{"A", "B"}, {"C"}, {"B", "C"}, {"D"},
	}, [][]int{{1, 2}, {0, 3}, {0, 3}, {1, 2}}, 0, 2))
}

func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	work := [][2]int{[2]int{id, 0}}
	seen := map[int]bool{}
	counts := map[string]int{}
	for len(work) > 0 {
		x := work[0]
		work = work[1:]
		if seen[x[0]] {
			continue
		}
		seen[x[0]] = true
		if x[1] == level {
			for _, vid := range watchedVideos[x[0]] {
				counts[vid]++
			}
			continue
		}

		for _, f := range friends[x[0]] {
			work = append(work, [2]int{f, x[1] + 1})
		}
	}

	var res []string
	for k := range counts {
		res = append(res, k)
	}
	sort.Slice(res, func(i, j int) bool {
		if counts[res[i]] == counts[res[j]] {
			return res[i] < res[j]
		}
		return counts[res[i]] < counts[res[j]]
	})
	return res
}
