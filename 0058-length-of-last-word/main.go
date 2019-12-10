package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("Hello"))
	fmt.Println(lengthOfLastWord(""))
	fmt.Println(lengthOfLastWord(" a "))
}

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	return len(s) - strings.LastIndexByte(s, ' ') - 1
}
