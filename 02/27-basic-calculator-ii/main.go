package main

import "fmt"

func main() {
	fmt.Println(calculate("3+2*2"))
	fmt.Println(calculate(" 3/2 "))
	fmt.Println(calculate(" 3+5 / 2 "))
}

func calculate(s string) int {
	var nums []int
	var ops []byte

	for len(s) > 0 {
		if s[0] == ' ' {
			s = s[1:]
			continue
		}
		if s[0] == '+' || s[0] == '-' || s[0] == '*' || s[0] == '/' {
			ops = append(ops, s[0])
			s = s[1:]
			continue
		}
		v := 0
		for len(s) > 0 && s[0] >= '0' && s[0] <= '9' {
			v *= 10
			v += int(s[0] - '0')
			s = s[1:]
		}
		nums = append(nums, v)
	}

	var anums []int
	var aops []byte
	curVal := nums[0]
	for i, n := range nums[1:] {
		op := ops[i]
		switch op {
		case '*':
			curVal *= n
		case '/':
			curVal /= n
		default:
			anums = append(anums, curVal)
			aops = append(aops, op)
			curVal = n
		}
	}
	anums = append(anums, curVal)

	curVal = anums[0]
	for i, n := range anums[1:] {
		op := aops[i]
		switch op {
		case '+':
			curVal += n
		case '-':
			curVal -= n
		}
	}

	return curVal
}
