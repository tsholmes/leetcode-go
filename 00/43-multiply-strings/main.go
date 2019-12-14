package main

import "fmt"

func main() {
	fmt.Println(multiply("2", "3"))
	fmt.Println(multiply("123", "456"))
}

func multiply(num1 string, num2 string) string {
	res := make([]byte, len(num1)+len(num2))
	var carry byte
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			k := len(res) - (len(num1) - i) - (len(num2) - j) + 1
			v := (num1[i]-'0')*(num2[j]-'0') + carry + res[k]
			res[k] = v % 10
			carry = v / 10
		}
		k := len(res) - (len(num1) - i) - len(num2)
		for carry > 0 {
			v := res[k] + carry
			res[k] = v % 10
			carry = v / 10
			k--
		}
	}
	for len(res) > 1 && res[0] == 0 {
		res = res[1:]
	}
	for i := range res {
		res[i] += '0'
	}
	return string(res)
}
