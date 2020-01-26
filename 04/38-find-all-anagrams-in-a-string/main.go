package main

import "fmt"

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("cbaebabacb", "abc"))
}

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}
	var ex [26]int
	for i := 0; i < len(p); i++ {
		ex[p[i]-'a']++
	}

	var i, j int
	var counts [26]int
	for j < len(p) {
		counts[s[j]-'a']++
		j++
	}

	var res []int
	if counts == ex {
		res = append(res, 0)
	}
	for j < len(s) {
		counts[s[j]-'a']++
		counts[s[i]-'a']--
		i++
		j++
		if counts == ex {
			res = append(res, i)
		}
	}

	return res
}
