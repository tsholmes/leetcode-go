package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	l := 0
	for {
		match := l < len(strs[0])
		for _, s := range strs[1:] {
			match = match && l < len(s) && strs[0][l] == s[l]
		}
		if match {
			l++
		} else {
			break
		}
	}
	return strs[0][:l]
}
