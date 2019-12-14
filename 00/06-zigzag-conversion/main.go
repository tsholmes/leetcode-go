package main

import "fmt"

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	plen := (numRows - 1) * 2
	res := make([]byte, 0, len(s))

	for line := 0; line < numRows; line++ {
		for i := 0; i < len(s); i += plen {
			l := i + line
			r := i + plen - line
			if l < len(s) {
				res = append(res, s[l])
			}
			if r != l && line != 0 && r < len(s) {
				res = append(res, s[r])
			}
		}
	}

	return string(res)
}
