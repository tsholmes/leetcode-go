package main

import "fmt"

func main() {
	fmt.Println(findSubstringInWraproundString("a"))
	fmt.Println(findSubstringInWraproundString("cac"))
	fmt.Println(findSubstringInWraproundString("zab"))
	fmt.Println(findSubstringInWraproundString("aabb"))
}

func findSubstringInWraproundString(p string) int {
	var longest [26]int

	for i := byte(0); i < 26; i++ {
		curLen := 0
		for j := 0; j < len(p); j++ {
			if p[j]-'a' != byte((int(i)+curLen)%26) {
				curLen = 0
			}
			if p[j]-'a' == byte((int(i)+curLen)%26) {
				curLen++
			}
			if curLen > longest[i] {
				longest[i] = curLen
			}
		}
	}

	res := 0
	for _, c := range longest[:] {
		res += c
	}

	return res
}
