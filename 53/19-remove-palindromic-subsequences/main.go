package main

import "fmt"

func main() {
	fmt.Println(removePalindromeSub("baabb"))
	fmt.Println(removePalindromeSub("bbabb"))
	fmt.Println(removePalindromeSub("bbaabaaa"))
}

func removePalindromeSub(s string) int {
	if s == "" {
		return 0
	}
	isp := true
	for i := 0; i < len(s)/2; i++ {
		isp = isp && s[i] == s[len(s)-i-1]
	}
	if isp {
		return 1
	}
	return 2
}
