package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	groups := map[string][]string{}
	var bs []byte
	for _, s := range strs {
		bs = append(bs[:0], s...)
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})
		groups[string(bs)] = append(groups[string(bs)], s)
	}

	var res [][]string
	for _, v := range groups {
		res = append(res, v)
	}

	return res
}
