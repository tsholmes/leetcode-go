package main

import "strings"

func main() {

}

func addBoldTag(s string, dict []string) string {
	open := false
	var res []byte

	rem := 0

	for len(s) > 0 {
		for _, s2 := range dict {
			if len(s2) > rem && strings.HasPrefix(s, s2) {
				rem = len(s2)
			}
		}
		if rem > 0 && !open {
			open = true
			res = append(res, '<', 'b', '>')
		} else if rem <= 0 && open {
			open = false
			res = append(res, '<', '/', 'b', '>')
		}
		res = append(res, s[0])
		s = s[1:]
		rem--
	}

	if open {
		res = append(res, '<', '/', 'b', '>')
	}
	return string(res)
}
