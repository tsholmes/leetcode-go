package main

import "fmt"

func main() {

}

func getHint(secret string, guess string) string {
	var counts [10]int
	var a, b int
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			a++
		} else {
			counts[secret[i]-'0']++
		}
	}
	for i := 0; i < len(guess); i++ {
		k := guess[i] - '0'
		if secret[i] != guess[i] && counts[k] > 0 {
			counts[k]--
			b++
		}
	}
	return fmt.Sprintf("%dA%dB", a, b)
}
