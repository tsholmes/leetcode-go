package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fractionToDecimal(1, 2))
	fmt.Println(fractionToDecimal(2, 1))
	fmt.Println(fractionToDecimal(2, 3))
	fmt.Println(fractionToDecimal(1, 7))
	fmt.Println(fractionToDecimal(22, 7))
	fmt.Println(fractionToDecimal(-22, 7))
	fmt.Println(fractionToDecimal(-22, -7))
	fmt.Println(fractionToDecimal(22, -7))
	fmt.Println(fractionToDecimal(0, 1))
}

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	neg := false
	if numerator < 0 {
		numerator = -numerator
		neg = !neg
	}
	if denominator < 0 {
		denominator = -denominator
		neg = !neg
	}
	var res []byte
	if neg {
		res = append(res, '-')
	}
	res = strconv.AppendInt(res, int64(numerator/denominator), 10)
	numerator = numerator % denominator
	if numerator == 0 {
		return string(res)
	}
	res = append(res, '.')
	seen := map[int]int{}

	for numerator > 0 {
		if i, ok := seen[numerator]; ok {
			res = append(res, 0)
			copy(res[i+1:], res[i:])
			res[i] = '('
			res = append(res, ')')
			break
		}
		seen[numerator] = len(res)
		numerator *= 10
		d := numerator / denominator
		numerator = numerator % denominator
		res = append(res, byte(d)+'0')
	}
	return string(res)
}
