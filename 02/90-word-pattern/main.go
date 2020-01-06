package main

import "strings"

func main() {

}

func wordPattern(pattern string, str string) bool {
	ws := strings.Split(str, " ")
	var cws [26]string
	used := map[string]bool{}
	if len(pattern) != len(ws) {
		return false
	}
	for i, w := range ws {
		k := pattern[i] - 'a'
		if cws[k] != "" {
			if w != cws[k] {
				return false
			}
		} else {
			if used[w] {
				return false
			}
			cws[k] = w
			used[w] = true
		}
	}
	return true
}
