package main

import "fmt"

func main() {
	fmt.Println(getPermutation(3, 3))
	fmt.Println(getPermutation(4, 9))
	// for i := 1; i <= 6; i++ {
	// 	fmt.Println(getPermutation(3, i))
	// }
}

func getPermutation(n int, k int) string {
	used := make([]bool, n)
	res := make([]byte, n)
	f := 1
	for i := 1; i <= n; i++ {
		f *= i
	}
	k--
	for i := 0; i < n; i++ {
		f /= (n - i)
		p := k / f
		k = k % f
		for j := 0; j <= p; j++ {
			if used[j] {
				p++
			}
		}
		res[i] = byte(p) + '1'
		used[p] = true
	}
	return string(res)
}
