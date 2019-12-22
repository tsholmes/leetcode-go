package main

import "fmt"

func main() {
	// for i := 1; i <= 28*26; i++ {
	// 	fmt.Println(convertToTitle(i))
	// }
	fmt.Println(convertToTitle(4))
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(701))
	fmt.Println(convertToTitle(3234))
}

func convertToTitle(n int) string {
	n--
	mul := 26
	for n >= mul {
		n -= mul
		mul *= 26
	}
	mul /= 26
	var res []byte
	for mul > 0 {
		res = append(res, byte(n/mul)+'A')
		n = n % mul
		mul /= 26
	}
	return string(res)
}
