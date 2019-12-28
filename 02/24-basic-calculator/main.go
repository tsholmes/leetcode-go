package main

import "fmt"

func main() {
	fmt.Println(calculate("1 + 1"))
	fmt.Println(calculate("2-1 + 2"))
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
}

func calculate(s string) int {
	neg := false
	negStack := []bool{false, false}
	val := 0

	for len(s) > 0 {
		switch s[0] {
		case ' ':
			s = s[1:]
		case '-':
			neg = !neg
			negStack = append(negStack, neg)
			s = s[1:]
		case '+':
			negStack = append(negStack, neg)
			s = s[1:]
		case '(':
			negStack = append(negStack, neg)
			s = s[1:]
		case ')':
			negStack = negStack[:len(negStack)-1]
			neg = negStack[len(negStack)-1]
			s = s[1:]
		default:
			v := 0
			for len(s) > 0 && s[0] >= '0' && s[0] <= '9' {
				v *= 10
				v += int(s[0] - '0')
				s = s[1:]
			}
			if neg {
				v = -v
			}
			val += v
			negStack = negStack[:len(negStack)-1]
			neg = negStack[len(negStack)-1]
		}
	}

	return val
}
