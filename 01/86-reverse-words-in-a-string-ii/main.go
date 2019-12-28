package main

import "fmt"

func main() {
	do := func(s string) {
		b := []byte(s)
		reverseWords(b)
		fmt.Println(string(b))
	}
	do("the sky is blue")
	do("abcdefg")
}

func reverseWords(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}

	i := 0
	for i < len(s) {
		j := i + 1
		for j < len(s) && s[j] != ' ' {
			j++
		}
		w := s[i:j]
		for k := 0; k < len(w)/2; k++ {
			l := len(w) - k - 1
			w[k], w[l] = w[l], w[k]
		}
		i = j + 1
	}
}
