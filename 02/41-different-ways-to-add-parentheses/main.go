package main

import "fmt"

func main() {
	fmt.Println(diffWaysToCompute("2-1-1"))
	fmt.Println(diffWaysToCompute("2*3-4*5"))
	fmt.Println(diffWaysToCompute("1"))
	fmt.Println(diffWaysToCompute("1+1"))
}

func diffWaysToCompute(input string) []int {
	var nums []int
	var ops []byte
	i := 0
	for i < len(input) {
		c := input[i]
		if c == '*' || c == '+' || c == '-' {
			ops = append(ops, c)
			i++
			continue
		}
		v := 0
		for i < len(input) && input[i] >= '0' && input[i] <= '9' {
			v *= 10
			v += int(input[i] - '0')
			i++
		}
		nums = append(nums, v)
	}

	calc := func(a, b int, op byte) int {
		switch op {
		case '*':
			return a * b
		case '+':
			return a + b
		default:
			return a - b
		}
	}

	var search func(int, int) []int
	search = func(i, j int) []int {
		if i == j {
			return nums[i : i+1]
		}
		if j == i+1 {
			return []int{calc(nums[i], nums[j], ops[i])}
		}

		var res []int

		for k := i; k < j; k++ {
			l, r := search(i, k), search(k+1, j)
			for _, lv := range l {
				for _, rv := range r {
					res = append(res, calc(lv, rv, ops[k]))
				}
			}
		}

		return res
	}

	return search(0, len(nums)-1)
}
