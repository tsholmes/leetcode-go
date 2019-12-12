package main

import "fmt"

func main() {
	fmt.Println(isScramble("great", "rgeat"))
	fmt.Println(isScramble("abcde", "caebd"))
	fmt.Println(isScramble("a", "a"))
	fmt.Println(isScramble("a", "b"))
}

func isScramble(s1 string, s2 string) bool {
	if len(s1) == 0 {
		return false
	}
	memo := map[[3]int]bool{}
	var search func(int, int, int) bool
	search = func(i, j, k int) bool {
		l := j - i
		if l == 1 {
			return s1[i] == s2[k]
		}
		key := [3]int{i, j, k}
		v, ok := memo[key]
		if ok {
			return v
		}

		for split := 1; split < l; split++ {
			s2 := l - split
			if (search(i, i+split, k) && search(i+split, j, k+split)) || (search(i, i+split, k+s2) && search(i+split, j, k)) {
				v = true
				break
			}
		}
		memo[key] = v
		return v
	}
	return search(0, len(s1), 0)
}
