package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fizzBuzz(15))
}

func fizzBuzz(n int) (res []string) {
	defer func() {
		recover()
	}()

	var do func(int)

	fb := [4]func(int) string{
		func(i int) string { return strconv.Itoa(i) },
		func(int) string { return "Fizz" },
		func(int) string { return "Buzz" },
		func(int) string { return "FizzBuzz" },
	}

	zero := func(v int) int {
		if v == 0 {
			return 1
		}
		return 0
	}
	do = func(i int) {
		res = append(res, fb[zero(i%3)+zero(i%5)<<1+zero(n-i+1)<<2](i))
		do(i + 1)
	}

	do(1)

	return nil
}
