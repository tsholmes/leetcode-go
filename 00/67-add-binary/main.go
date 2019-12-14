package main

import "fmt"

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
	fmt.Println(addBinary("0", "0"))
	fmt.Println(addBinary("0", "1"))
}

func addBinary(a string, b string) string {
	l := len(a)
	if len(b) > l {
		l = len(b)
	}
	res := make([]byte, l+1)
	carry := 0
	for i := 0; i <= l; i++ {
		if i < len(a) {
			carry += int(a[len(a)-i-1] - '0')
		}
		if i < len(b) {
			carry += int(b[len(b)-i-1] - '0')
		}
		res[l-i] = byte(carry&1) + '0'
		carry /= 2
	}
	if res[0] == '0' {
		res = res[1:]
	}
	return string(res)
}
