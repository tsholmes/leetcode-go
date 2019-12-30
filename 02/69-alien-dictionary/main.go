package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(alienOrder([]string{"wrt", "wrf", "er", "ett", "rftt"}))
	fmt.Println(alienOrder([]string{"z", "x"}))
	fmt.Println(alienOrder([]string{"z", "x", "z"}))
	fmt.Println(alienOrder([]string{"a", "b", "ca", "cc"}))
	fmt.Println(alienOrder([]string{"wnlb"}))
}

func alienOrder(words []string) string {
	if len(words) == 1 {
		return words[0]
	}
	edges := map[byte]map[byte]struct{}{}
	cs := map[byte]struct{}{}
	for i := 1; i < len(words); i++ {
		a, b := words[i-1], words[i]
		for j := 0; j < len(a) && j < len(b); j++ {
			if a[j] != b[j] {
				if _, ok := edges[a[j]]; !ok {
					edges[a[j]] = map[byte]struct{}{}
				}
				edges[a[j]][b[j]] = struct{}{}
				break
			}
		}
		for _, c := range []byte(a) {
			cs[c] = struct{}{}
		}
		for _, c := range []byte(b) {
			cs[c] = struct{}{}
		}
	}
	depth := map[byte]int{}
	var search func(byte, int, map[byte]bool)
	valid := true
	search = func(c byte, d int, seen map[byte]bool) {
		if seen[c] {
			valid = false
			return
		}
		seen[c] = true
		if d >= depth[c] {
			depth[c] = d
		}
		for c2 := range edges[c] {
			search(c2, d+1, seen)
		}
		seen[c] = false
	}

	res := make([]byte, 0, len(cs))
	for c := range cs {
		res = append(res, c)
		if _, ok := depth[c]; ok {
			continue
		}
		search(c, 0, map[byte]bool{})
	}
	if !valid {
		return ""
	}

	sort.Slice(res, func(i, j int) bool { return depth[res[i]] < depth[res[j]] })

	return string(res)
}
