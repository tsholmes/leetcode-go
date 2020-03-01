package main

import "fmt"

func main() {
	fmt.Println(isSolvable([]string{"SEND", "MORE"}, "MONEY"))
	fmt.Println(isSolvable([]string{"SIX", "SEVEN", "SEVEN"}, "TWENTY"))
	fmt.Println(isSolvable([]string{"THIS", "IS", "TOO"}, "FUNNY"))
	fmt.Println(isSolvable([]string{"LEET", "CODE"}, "POINT"))
}

func isSolvable(words []string, result string) bool {
	for i := range words {
		if len(words[i]) > len(result) {
			return false
		}
	}

	cs := make([][]byte, len(result))
	rs := make([]byte, len(result))
	for i := 0; i < len(result); i++ {
		for _, w := range words {
			j := len(w) - i - 1
			if j >= 0 {
				cs[i] = append(cs[i], w[j])
			}
		}
		rs[len(result)-i-1] = result[i]
	}

	var search func(int, int, int) bool
	ds := map[byte]int{}
	used := make([]bool, 10)
	search = func(i, j, carry int) bool {
		if i == len(result) {
			return carry == 0
		}
		if j == len(cs[i]) {
			s := carry
			for _, c := range cs[i] {
				s += ds[c]
			}
			s, carry = s%10, s/10
			if d, ok := ds[rs[i]]; ok {
				if d != s {
					return false
				}
				return search(i+1, 0, carry)
			} else if !used[s] {
				used[s] = true
				ds[rs[i]] = s
				defer func() {
					used[s] = false
					delete(ds, rs[i])
				}()
				return search(i+1, 0, carry)
			} else {
				return false
			}
		}
		c := cs[i][j]
		if _, ok := ds[c]; ok {
			return search(i, j+1, carry)
		}
		for d := 0; d < 10; d++ {
			if used[d] {
				continue
			}
			used[d] = true
			ds[c] = d
			res := search(i, j+1, carry)
			used[d] = false
			delete(ds, c)
			if res {
				return true
			}
		}
		return false
	}
	return search(0, 0, 0)
}
