package main

import "fmt"

func main() {
	fmt.Println(isAdditiveNumber("112358"))
	fmt.Println(isAdditiveNumber("199100199"))
	fmt.Println(isAdditiveNumber("1"))
	fmt.Println(isAdditiveNumber("000"))
	fmt.Println(isAdditiveNumber("199111992"))
	fmt.Println(isAdditiveNumber("1991000199299498797"))
}

func isAdditiveNumber(num string) bool {
	maxLen := len(num)

	digits := make([]byte, maxLen)

	for l1 := 1; l1 <= maxLen; l1++ {
		if num[0] == '0' && l1 > 1 {
			continue
		}
		n1 := 0
		for i := 0; i < l1; i++ {
			n1 *= 10
			n1 += int(num[i] - '0')
		}
		for l2 := 1; l2 <= maxLen; l2++ {
			if l1+l2 >= len(num) {
				continue
			}
			if num[l1] == '0' && l2 > 1 {
				continue
			}
			n2 := 0
			for i := 0; i < l2; i++ {
				n2 *= 10
				n2 += int(num[l1+i] - '0')
			}

			i := l1 + l2
			valid := true
			a, b := n1, n2
			for i < len(num) && valid {
				c := a + b
				if c == 0 {
					valid = num[i] == '0'
					i++
					continue
				}
				dc := 0
				cc := c
				for cc > 0 {
					dc++
					digits[maxLen-dc] = byte(cc%10) + '0'
					cc /= 10
				}
				if i+dc > len(num) {
					valid = false
					break
				}
				if num[i:i+dc] != string(digits[maxLen-dc:]) {
					valid = false
					break
				}
				a, b = b, c
				i += dc
			}
			if valid {
				return true
			}
		}
	}

	return false
}
