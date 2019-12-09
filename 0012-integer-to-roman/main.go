package main

import "fmt"

func main() {
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(777))
	fmt.Println(intToRoman(444))
	fmt.Println(intToRoman(1994))
}

func intToRoman(num int) string {
	var res []byte

	// 1000s
	for num >= 1000 {
		res = append(res, 'M')
		num -= 1000
	}

	// 100s
	if num >= 900 {
		res = append(res, 'C', 'M')
		num -= 900
	}
	if num >= 500 {
		res = append(res, 'D')
		num -= 500
	}
	if num >= 400 {
		res = append(res, 'C', 'D')
		num -= 400
	}
	for num >= 100 {
		res = append(res, 'C')
		num -= 100
	}

	// 10s
	if num >= 90 {
		res = append(res, 'X', 'C')
		num -= 90
	}
	if num >= 50 {
		res = append(res, 'L')
		num -= 50
	}
	if num >= 40 {
		res = append(res, 'X', 'L')
		num -= 40
	}
	for num >= 10 {
		res = append(res, 'X')
		num -= 10
	}

	// 1s
	if num == 9 {
		res = append(res, 'I', 'X')
		num -= 9
	}
	if num >= 5 {
		res = append(res, 'V')
		num -= 5
	}
	if num == 4 {
		res = append(res, 'I', 'V')
		num -= 4
	}
	for num > 0 {
		res = append(res, 'I')
		num--
	}

	return string(res)
}
