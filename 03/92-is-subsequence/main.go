package main

import "strings"

func main() {

}

func isSubsequence(s string, t string) bool {
	for len(s) > 0 {
		i := strings.IndexByte(t, s[0])
		if i == -1 {
			return false
		}
		s = s[1:]
		t = t[i+1:]
	}
	return true
}
