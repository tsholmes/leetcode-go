package main

import "fmt"

func main() {
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3, 4}))
	fmt.Println(numberOfArithmeticSlices([]int{1, 3, 5, 7, 9}))
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3, 5, 7, 10, 13}))
	fmt.Println(numberOfArithmeticSlices([]int{1, 2}))
}

func numberOfArithmeticSlices(A []int) int {
	res := 0
	i := 0
	for i < len(A)-2 {
		d := A[i+1] - A[i]
		j := i + 2
		for j < len(A) && A[j]-A[j-1] == d {
			j++
		}
		res += (j - i - 2) * (j - i - 1) / 2
		i = j - 1
	}
	return res
}
