package main

import "fmt"

func main() {
	fmt.Println(maxScore("011101"))
	fmt.Println(maxScore("00111"))
	fmt.Println(maxScore("1111"))
}

func maxScore(s string) int {
	t1s := 0
	for _, c := range []byte(s) {
		if c == '1' {
			t1s++
		}
	}

	max := 0
	cur1s := 0
	cur0s := 0
	for _, c := range []byte(s[:len(s)-1]) {
		if c == '0' {
			cur0s++
		} else {
			cur1s++
		}
		if cur0s+t1s-cur1s > max {
			max = cur0s + t1s - cur1s
		}
	}
	return max
}
