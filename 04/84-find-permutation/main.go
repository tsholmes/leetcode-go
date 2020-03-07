package main

func main() {

}

func findPermutation(s string) []int {
	s += "I"
	var res []int
	next := 1
	for len(s) > 0 {
		if s[0] == 'I' {
			res = append(res, next)
			next++
			s = s[1:]
			continue
		}
		c := 1
		for s[0] == 'D' {
			c++
			s = s[1:]
		}
		s = s[1:]
		for j := c - 1; j >= 0; j-- {
			res = append(res, next+j)
		}
		next += c
	}
	return res
}
