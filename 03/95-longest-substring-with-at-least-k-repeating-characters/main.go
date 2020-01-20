package main

import "fmt"

func main() {
	fmt.Println(longestSubstring("aaabb", 3))
}

func longestSubstring(s string, k int) int {
	max := 0
	for i := 0; i < len(s); i++ {
		var counts [26]int
		for j := i + 1; j <= len(s); j++ {
			counts[s[j-1]-'a']++
			if j-i <= max {
				continue
			}
			valid := true
			for _, c := range counts {
				if c > 0 && c < k {
					valid = false
					break
				}
			}
			if valid {
				max = j - i
			}
		}
	}
	return max
}
