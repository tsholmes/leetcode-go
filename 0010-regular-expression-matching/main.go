package main

import "fmt"

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
}

func isMatch(s string, p string) bool {
	type match struct {
		ch  byte
		any bool
		rep bool
		fin bool

		next []int
	}

	ms := []match{}
	for i := 0; i < len(p); i++ {
		c := p[i]
		rep := false
		any := c == '.'
		if i+1 < len(p) && p[i+1] == '*' {
			rep = true
			i++
		}
		ms = append(ms, match{ch: c, rep: rep, any: any})
	}
	ms = append(ms, match{fin: true})

	for i := range ms[:len(ms)-1] {
		if ms[i].rep {
			ms[i].next = append(ms[i].next, i)
		}
		j := i + 1
		for j < len(ms) {
			ms[i].next = append(ms[i].next, j)
			if !ms[j].rep {
				break
			}
			j++
		}
	}

	state := []int{0}
	for i, m := range ms[:len(ms)-1] {
		if m.rep {
			state = append(state, i+1)
		} else {
			break
		}
	}

	for _, b := range []byte(s) {
		next := []int{}
		seen := map[int]bool{}

		for _, j := range state {
			m := ms[j]
			if m.any || b == m.ch {
				for _, k := range m.next {
					if !seen[k] {
						seen[k] = true
						next = append(next, k)
					}
				}
			}
		}
		state = next
	}

	fin := false
	for _, i := range state {
		fin = fin || ms[i].fin
	}

	return fin
}
