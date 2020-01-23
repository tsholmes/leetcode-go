package main

func main() {

}

func longestPalindrome(s string) int {
	var counts [256]int
	for i := 0; i < len(s); i++ {
		counts[s[i]]++
	}
	odd := false
	evens := 0
	for _, c := range counts {
		if c&1 == 1 {
			odd = true
		}
		evens += c &^ 1
	}
	if odd {
		evens++
	}
	return evens
}
