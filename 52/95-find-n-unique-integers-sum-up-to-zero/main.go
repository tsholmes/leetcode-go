package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(sumZero(i))
	}
}

func sumZero(n int) []int {
	if n == 0 {
		return nil
	}
	var res []int
	for i := 1; i <= n/2; i++ {
		res = append(res, i, -i)
	}
	if n%2 == 1 {
		res = append(res, 0)
	}
	return res
}
