package main

import "sort"

func main() {

}

func frequencySort(s string) string {
	var counts [256]int
	for i := 0; i < len(s); i++ {
		counts[s[i]]++
	}
	bs := []byte(s)
	sort.Slice(bs, func(i, j int) bool {
		ci, cj := counts[bs[i]], counts[bs[j]]
		if ci == cj {
			return bs[i] < bs[j]
		}
		return ci > cj
	})
	return string(bs)
}
