package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}

func evalRPN(tokens []string) int {
	var stack []int
	take := func() (int, int) {
		a, b := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		return a, b
	}
	for _, t := range tokens {
		var a, b int
		switch t {
		case "+":
			a, b = take()
			stack = append(stack, a+b)
		case "-":
			a, b = take()
			stack = append(stack, a-b)
		case "*":
			a, b = take()
			stack = append(stack, a*b)
		case "/":
			a, b = take()
			stack = append(stack, a/b)
		default:
			a, _ = strconv.Atoi(t)
			stack = append(stack, a)
		}
	}
	return stack[0]
}
