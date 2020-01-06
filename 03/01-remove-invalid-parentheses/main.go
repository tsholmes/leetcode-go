package main

import "fmt"

func main() {
	fmt.Println(removeInvalidParentheses("()())()"))
	fmt.Println(removeInvalidParentheses("(a)())()"))
	fmt.Println(removeInvalidParentheses(")("))
}

func removeInvalidParentheses(s string) []string {
	var res []string
	work := make([]byte, 0, len(s))
	seen := map[string]bool{}
	var search func(int, int, int)
	search = func(i int, rem int, open int) {
		if len(s)-i < open {
			return
		}
		if len(res) > 0 && len(s)-len(res[0]) < rem {
			return
		}
		if i == len(s) {
			if len(res) > 0 && len(work) > len(res[0]) {
				res = res[:0]
				seen = map[string]bool{}
			}
			s := string(work)
			if !seen[s] {
				res = append(res, s)
				seen[s] = true
			}
			return
		}
		wl := len(work)
		switch s[i] {
		case '(':
			search(i+1, rem+1, open)
			work = append(work, s[i])
			search(i+1, rem, open+1)
			work = work[:wl]
		case ')':
			search(i+1, rem+1, open)
			if open > 0 {
				work = append(work, s[i])
				search(i+1, rem, open-1)
				work = work[:wl]
			}
		default:
			work = append(work, s[i])
			search(i+1, rem, open)
			work = work[:wl]
		}
	}
	search(0, 0, 0)
	return res
}
