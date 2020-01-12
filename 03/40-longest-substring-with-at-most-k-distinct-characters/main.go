package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstringKDistinct("eceba", 2))
	fmt.Println(lengthOfLongestSubstringKDistinct("aa", 1))
}

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if k == 0 {
		return 0
	}
	var counts [256]int
	var distinct int
	max := 0
	j := 0
	for i := 0; i < len(s); i++ {
		counts[s[i]]++
		if counts[s[i]] == 1 {
			distinct++
		}
		for distinct > k {
			counts[s[j]]--
			if counts[s[j]] == 0 {
				distinct--
			}
			j++
		}
		if i-j+1 > max {
			max = i - j + 1
		}
	}
	return max
}
