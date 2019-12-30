package main

import "fmt"

func main() {
	fmt.Println(groupStrings([]string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"}))
}

func groupStrings(strings []string) [][]string {
	groups := map[string][]string{}

	for _, s := range strings {
		b := []byte(s)
		for i := range b {
			b[i] = ((b[i] - s[0] + 26) % 26) + 'a'
		}
		groups[string(b)] = append(groups[string(b)], s)
	}

	var res [][]string
	for _, v := range groups {
		res = append(res, v)
	}
	return res
}
