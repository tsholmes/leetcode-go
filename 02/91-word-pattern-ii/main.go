package main

import "fmt"

func main() {
	fmt.Println(wordPatternMatch("abab", "redblueredblue"))
	fmt.Println(wordPatternMatch("aaaa", "asdasdasdasd"))
	fmt.Println(wordPatternMatch("aabb", "xyzabcxzyabc"))
}

func wordPatternMatch(pattern string, str string) bool {
	used := map[string]bool{}
	var wcs [26]string
	var search func(i, j int) bool
	search = func(i, j int) bool {
		if i == len(pattern) || j == len(str) {
			return i == len(pattern) && j == len(str)
		}
		k := int(pattern[i] - 'a')
		if wcs[k] != "" {
			for m := 0; m < len(wcs[k]); m++ {
				if j+m >= len(str) {
					return false
				}
				if wcs[k][m] != str[j+m] {
					return false
				}
			}
			return search(i+1, j+len(wcs[k]))
		}
		for m := 1; j+m <= len(str); m++ {
			s := str[j : j+m]
			if used[s] {
				continue
			}
			wcs[k] = s
			used[s] = true
			if search(i+1, j+m) {
				return true
			}
			wcs[k] = ""
			used[s] = false
		}
		return false
	}
	return search(0, 0)
}
