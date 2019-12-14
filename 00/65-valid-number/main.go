package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isNumber("0"))
	fmt.Println(isNumber(" 0.1 "))
	fmt.Println(isNumber("abc"))
	fmt.Println(isNumber("1 a"))
	fmt.Println(isNumber("2e10"))
	fmt.Println(isNumber(" -90e3   "))
	fmt.Println(isNumber(" 1e"))
	fmt.Println(isNumber("e3"))
	fmt.Println(isNumber(" 6e-1"))
	fmt.Println(isNumber(" 99e2.5 "))
	fmt.Println(isNumber("53.5e93"))
	fmt.Println(isNumber(" --6 "))
	fmt.Println(isNumber("-+3"))
	fmt.Println(isNumber("95a54e53"))
}

func isNumber(s string) bool {
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return false
	}

	if s[0] == '+' || s[0] == '-' {
		s = s[1:]
	}
	if len(s) == 0 {
		return false
	}

	digits := 0

	for len(s) > 0 && s[0] >= '0' && s[0] <= '9' {
		s = s[1:]
		digits++
	}

	if len(s) > 0 && s[0] == '.' {
		s = s[1:]
		for len(s) > 0 && s[0] >= '0' && s[0] <= '9' {
			s = s[1:]
			digits++
		}
	}

	if digits == 0 {
		return false
	}

	if len(s) > 0 && s[0] == 'e' {
		s = s[1:]
		if len(s) == 0 {
			return false
		}
		if s[0] == '+' || s[0] == '-' {
			s = s[1:]
		}
		if len(s) == 0 {
			return false
		}
		for len(s) > 0 && s[0] >= '0' && s[0] <= '9' {
			s = s[1:]
		}
	}

	return len(s) == 0
}
