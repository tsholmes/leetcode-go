package main

import "strings"

func main() {

}

func reverseWords(s string) string {
	ws := strings.Split(s, " ")
	for i, w := range ws {
		ws[i] = reverse(w)
	}
	return strings.Join(ws, " ")
}

func reverse(s string) string {
	res := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		res = append(res, s[len(s)-i-1])
	}
	return string(res)
}
