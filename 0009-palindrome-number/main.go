package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(1221))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	low := 1
	high := 1
	for x/(high*10) != 0 {
		high *= 10
	}

	for low < high {
		if (x/low)%10 != (x/high)%10 {
			return false
		}
		low *= 10
		high /= 10
	}

	return true
}
