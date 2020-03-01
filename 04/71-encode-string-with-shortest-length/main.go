package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(encode("aaa"))
	fmt.Println(encode("aaaaa"))
	fmt.Println(encode("aaaaaaaaaa"))
	fmt.Println(encode("aabcaabcd"))
	fmt.Println(encode("abbbabbbcabbbabbbc"))
}

func encode(s string) string {
	memo := map[string]string{}
	var work []byte
	var search func(string) string
	search = func(str string) string {
		if len(str) < 5 {
			return str
		}
		if v, ok := memo[str]; ok {
			return v
		}

		shortest := str

		// check if str is fully repeatable
		for reps := 2; reps <= len(str); reps++ {
			if len(str)%reps != 0 {
				continue
			}
			off := len(str) / reps
			valid := true
			for i := off; i < len(str); i++ {
				valid = valid && str[i] == str[i%off]
			}
			if valid {
				ss := search(str[:off])
				work = strconv.AppendInt(work[:0], int64(reps), 10)
				work = append(work, '[')
				work = append(work, ss...)
				work = append(work, ']')
				if len(work) < len(shortest) {
					shortest = string(work)
				}
			}
		}

		// check all splits
		for i := 1; i < len(str); i++ {
			sa, sb := search(str[:i]), search(str[i:])
			if len(sa)+len(sb) < len(shortest) {
				shortest = sa + sb
			}
		}

		memo[str] = shortest
		return shortest
	}
	return search(s)
}
