package main

import "fmt"

func main() {
	fmt.Println(numberToWords(123))
	fmt.Println(numberToWords(12345))
	fmt.Println(numberToWords(1234567))
	fmt.Println(numberToWords(1234567891))
}

var (
	ones = []string{
		"Zero",
		"One",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
	}
	teens = []string{
		"Ten",
		"Eleven",
		"Twelve",
		"Thirteen",
		"Fourteen",
		"Fifteen",
		"Sixteen",
		"Seventeen",
		"Eighteen",
		"Nineteen",
	}
	tens = []string{
		"",
		"",
		"Twenty",
		"Thirty",
		"Forty",
		"Fifty",
		"Sixty",
		"Seventy",
		"Eighty",
		"Ninety",
	}
)

func numberToWords(num int) string {
	var res []byte

	if num < 10 {
		return ones[num]
	}

	hundos := func(n int) {
		if n >= 100 {
			d := n / 100
			res = append(res, ' ')
			res = append(res, ones[d]...)
			res = append(res, " Hundred"...)
			n = n % 100
		}
		if n >= 20 {
			d := n / 10
			res = append(res, ' ')
			res = append(res, tens[d]...)
			n = n % 10
		} else if n >= 10 {
			res = append(res, ' ')
			res = append(res, teens[n-10]...)
			n = 0
		}
		if n > 0 {
			res = append(res, ' ')
			res = append(res, ones[n]...)
		}
	}

	if num >= 1000000000 {
		hundos(num / 1000000000)
		res = append(res, " Billion"...)
		num = num % 1000000000
	}
	if num >= 1000000 {
		hundos(num / 1000000)
		res = append(res, " Million"...)
		num = num % 1000000
	}
	if num >= 1000 {
		hundos(num / 1000)
		res = append(res, " Thousand"...)
		num = num % 1000
	}
	hundos(num)

	return string(res[1:])
}
