package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
}

func isIsomorphic(s string, t string) bool {
	m := map[byte]byte{}
	used := map[byte]bool{}
	for i := 0; i < len(s); i++ {
		a, b := s[i], t[i]
		if c, ok := m[a]; ok && c != b {
			return false
		} else if !ok && used[b] {
			return false
		}
		m[a] = b
		used[b] = true
	}
	return true
}
