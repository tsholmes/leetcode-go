package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("DCCLXXVII"))
	fmt.Println(romanToInt("CDXLIV"))
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	v := 0
	for i := 0; i < len(s); i++ {
		var c2 byte
		if i+1 < len(s) {
			c2 = s[i+1]
		}
		switch s[i] {
		case 'M':
			v += 1000
		case 'D':
			v += 500
		case 'C':
			if c2 == 'M' {
				v += 900
				i++
			} else if c2 == 'D' {
				v += 400
				i++
			} else {
				v += 100
			}
		case 'L':
			v += 50
		case 'X':
			if c2 == 'C' {
				v += 90
				i++
			} else if c2 == 'L' {
				v += 40
				i++
			} else {
				v += 10
			}
		case 'V':
			v += 5
		case 'I':
			if c2 == 'X' {
				v += 9
				i++
			} else if c2 == 'V' {
				v += 4
				i++
			} else {
				v++
			}
		}
	}
	return v
}
