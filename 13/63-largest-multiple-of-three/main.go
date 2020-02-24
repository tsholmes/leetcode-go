package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestMultipleOfThree([]int{5, 8}))
}

func largestMultipleOfThree(digits []int) string {
	sort.Ints(digits)
	sum := 0
	for _, n := range digits {
		sum += n
	}
	var ones []int
	var twos []int
	for i, n := range digits {
		if n%3 == 1 {
			ones = append(ones, i)
		} else if n%3 == 2 {
			twos = append(twos, i)
		}
	}
	var rem []int
	if sum%3 == 1 {
		if len(ones) > 0 {
			rem = ones[:1]
		} else if len(twos) > 1 {
			rem = twos[:2]
		} else {
			return ""
		}
	} else if sum%3 == 2 {
		if len(twos) > 0 {
			rem = twos[:1]
		} else if len(ones) > 1 {
			rem = ones[:2]
		} else {
			return ""
		}
	}
	sort.Ints(rem)
	count := 0
	for i, n := range digits {
		if len(rem) > 0 && rem[0] == i {
			rem = rem[1:]
			continue
		}
		digits[count] = n
		count++
	}
	digits = digits[:count]
	var res []byte
	for i := len(digits) - 1; i >= 0; i-- {
		res = append(res, byte(digits[i])+'0')
	}
	if len(res) > 0 && res[0] == '0' {
		return "0"
	}
	return string(res)
}
