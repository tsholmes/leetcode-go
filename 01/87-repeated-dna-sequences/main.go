package main

import "fmt"

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
}

func findRepeatedDnaSequences(s string) []string {
	counts := map[string]int{}
	for i := 0; i <= len(s)-10; i++ {
		counts[s[i:i+10]]++
	}
	var res []string
	for k, v := range counts {
		if v > 1 {
			res = append(res, k)
		}
	}
	return res
}
