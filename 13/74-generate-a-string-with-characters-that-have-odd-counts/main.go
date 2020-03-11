package main

import "fmt"

func main() {
	fmt.Println(generateTheString(4))
	fmt.Println(generateTheString(5))
}

func generateTheString(n int) string {
	var res []byte
	if n%2 == 0 {
		for i := 0; i < n-1; i++ {
			res = append(res, 'a')
		}
		res = append(res, 'b')
	} else {
		for i := 0; i < n; i++ {
			res = append(res, 'a')
		}
	}
	return string(res)
}
