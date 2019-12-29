package main

import "fmt"

func main() {
	fmt.Println(countDigitOne(13))

	for i := 1; i <= 10000; i++ {
		a, b := countDigitOne(i), countDigitOneSlow(i)
		if a != b {
			fmt.Println(i, "=>", a, b)
		}
	}
}

func countDigitOne(n int) int {
	if n <= 0 {
		return 0
	}
	if n < 10 {
		return 1
	}
	mul := 1
	it := 1
	for n >= mul*10 {
		mul *= 10
		it++
	}
	left := n / mul
	right := n % mul

	rs := (mul / 10) * (it - 1)
	rd := countDigitOne(right)

	res := rs*left + rd
	if left > 1 {
		res += mul
	} else {
		res += 1 + right
	}

	return res
}

func countDigitOneSlow(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		j := i
		for j > 0 {
			d := j % 10
			j /= 10
			if d == 1 {
				count++
			}
		}
	}
	return count
}
