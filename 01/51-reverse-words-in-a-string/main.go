package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("  hello world!  "))
	fmt.Println(reverseWords("a good   example"))
}

func reverseWords(s string) string {
	ws := strings.Split(s, " ")
	count := 0
	for _, w := range ws {
		if w != "" {
			ws[count] = w
			count++
		}
	}
	ws = ws[:count]
	for i := 0; i < len(ws)/2; i++ {
		j := len(ws) - i - 1
		ws[i], ws[j] = ws[j], ws[i]
	}
	return strings.Join(ws, " ")
}
