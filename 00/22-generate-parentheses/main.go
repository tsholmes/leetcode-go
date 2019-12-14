package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}

	cur := make([]byte, n*2)
	var res []string

	var walk func(int, int, int)
	walk = func(pos int, open int, used int) {
		if pos == n*2 {
			res = append(res, string(cur))
			return
		}
		if used < n {
			cur[pos] = '('
			walk(pos+1, open+1, used+1)
		}
		if open > 0 {
			cur[pos] = ')'
			walk(pos+1, open-1, used)
		}
	}

	walk(0, 0, 0)

	return res
}
