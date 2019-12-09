package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	src := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var res []string

	cur := make([]byte, len(digits))

	var search func(int)
	search = func(i int) {
		if i == len(digits) {
			res = append(res, string(cur))
			return
		}
		s := src[digits[i]]
		for j := 0; j < len(s); j++ {
			cur[i] = s[j]
			search(i + 1)
		}
	}

	search(0)

	return res
}
