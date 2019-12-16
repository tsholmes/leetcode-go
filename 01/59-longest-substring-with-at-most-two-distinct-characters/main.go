package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstringTwoDistinct("eceba"))
	fmt.Println(lengthOfLongestSubstringTwoDistinct("ccaabbb"))
}

func lengthOfLongestSubstringTwoDistinct(s string) int {
	if len(s) == 0 {
		return 0
	}
	max := 1
	for i := 0; i < len(s); i++ {
		c1 := s[i]
		var c2 byte
		for j := i + 1; j < len(s); j++ {
			if s[j] == c1 || s[j] == c2 || c2 == 0 {
				if s[j] != c1 {
					c2 = s[j]
				}
				if (j - i + 1) > max {
					max = j - i + 1
				}
			} else {
				break
			}
		}
	}
	return max
}
