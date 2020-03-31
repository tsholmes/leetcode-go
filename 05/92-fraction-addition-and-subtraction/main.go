package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

}

func fractionAddition(expression string) string {
	if expression[0] != '-' {
		expression = "+" + expression
	}

	getNum := func() int {
		end := strings.IndexAny(expression, "/+-")
		if end == -1 {
			end = len(expression)
		}
		n, _ := strconv.Atoi(expression[:end])
		expression = expression[end:]
		return n
	}

	num, den := 0, 1
	for len(expression) > 0 {
		neg := expression[0] == '-'
		expression = expression[1:]

		n := getNum()
		expression = expression[1:]
		d := getNum()

		if neg {
			n = -n
		}

		num = num*d + n*den
		den = den * d

		if num != 0 {
			f := gcf(abs(num), den)
			num /= f
			den /= f
		} else {
			num, den = 0, 1
		}
	}

	return fmt.Sprintf("%d/%d", num, den)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func gcf(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
