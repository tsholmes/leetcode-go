package main

import "fmt"

func main() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
}

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return nil
	}

	var res []int

	sz := len(words[0])
	tlen := sz * len(words)

	for i := 0; i <= len(s)-tlen; i++ {
		used := make([]bool, len(words))
		valid := true
		for j := i; j < i+tlen; j += sz {
			found := false
			for k, w := range words {
				if used[k] {
					continue
				}
				if s[j:j+sz] == w {
					used[k] = true
					found = true
					break
				}
			}
			if !found {
				valid = false
				break
			}
		}
		if valid {
			res = append(res, i)
		}
	}

	return res
}
