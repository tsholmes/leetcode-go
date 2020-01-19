package main

func main() {

}

func firstUniqChar(s string) int {
	first := map[byte]int{}
	uniq := map[byte]bool{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if _, ok := uniq[c]; !ok {
			uniq[c] = true
			first[c] = i
		} else {
			uniq[c] = false
		}
	}
	min := len(s)
	for k, v := range uniq {
		if v && first[k] < min {
			min = first[k]
		}
	}
	if min == len(s) {
		return -1
	}
	return min
}
