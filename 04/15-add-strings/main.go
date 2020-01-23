package main

import "fmt"

func main() {
	fmt.Println(addStrings("1", "1"))
}

func addStrings(num1 string, num2 string) string {
	var res []byte
	carry := 0
	for i := 0; i < len(num1) || i < len(num2); i++ {
		d1, d2 := 0, 0
		if i < len(num1) {
			d1 = int(num1[len(num1)-i-1] - '0')
		}
		if i < len(num2) {
			d2 = int(num2[len(num2)-i-1] - '0')
		}
		d := d1 + d2 + carry
		d, carry = d%10, d/10
		res = append(res, byte(d)+'0')
	}
	if carry > 0 {
		res = append(res, byte(carry)+'0')
	}
	for i := 0; i < len(res)/2; i++ {
		j := len(res) - i - 1
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
