package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "bba"))
	fmt.Println(strStr("whatever", ""))
}

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
