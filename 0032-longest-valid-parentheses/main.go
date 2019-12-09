package main

import "fmt"

func main() {
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses("(()))((()))"))
	fmt.Println(longestValidParentheses("(())()(()(("))
}

func longestValidParentheses(s string) int {
	max := 0

	var opens []int
	depth := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			if len(opens) <= depth {
				opens = append(opens, i)
			}
			depth++
		} else if depth > 0 {
			depth--
			j := opens[depth]
			opens = opens[:depth+1]
			l := i - j + 1
			if l > max {
				max = l
			}
		} else {
			depth = 0
			opens = opens[:0]
		}
	}

	return max
}
