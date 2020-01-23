package main

import "fmt"

func main() {
	fmt.Println(validWordAbbreviation("internationalization", "i12iz4n"))
	fmt.Println(validWordAbbreviation("apple", "a2e"))
	fmt.Println(validWordAbbreviation("apple", "a4"))
}

func validWordAbbreviation(word string, abbr string) bool {
	for len(abbr) > 0 {
		if len(word) == 0 {
			return false
		}
		if abbr[0] >= '0' && abbr[0] <= '9' {
			c := 0
			for len(abbr) > 0 && abbr[0] >= '0' && abbr[0] <= '9' {
				if c == 0 && abbr[0] == '0' {
					return false
				}
				c = c*10 + int(abbr[0]-'0')
				abbr = abbr[1:]
			}
			if c > len(word) {
				return false
			}
			word = word[c:]
		} else {
			if abbr[0] != word[0] {
				return false
			}
			abbr = abbr[1:]
			word = word[1:]
		}
	}
	return len(word) == 0
}
