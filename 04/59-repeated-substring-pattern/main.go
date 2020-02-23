package main

func main() {

}

func repeatedSubstringPattern(s string) bool {
	for off := 1; off <= len(s)/2; off++ {
		if len(s)%off != 0 {
			continue
		}
		ok := true
		for i := 0; i < off && ok; i++ {
			for j := i + off; j < len(s) && ok; j += off {
				ok = ok && s[i] == s[j]
			}
		}
		if ok {
			return true
		}
	}
	return false
}
