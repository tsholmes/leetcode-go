package main

import "fmt"

func main() {
	fmt.Println(partition("aab"))
}

func partition(s string) [][]string {
	var res [][]string
	work := make([]string, 0, len(s))
	var search func(int)
	search = func(i int) {
		if i == len(s) {
			res = append(res, append([]string{}, work...))
			return
		}
		wl := len(work)
		for l := 1; i+l <= len(s); l++ {
			valid := true
			for j := 0; j < l/2; j++ {
				valid = valid && s[i+j] == s[i+(l-j-1)]
			}
			if !valid {
				continue
			}
			work = append(work, s[i:i+l])
			search(i + l)
			work = work[:wl]
		}
	}
	search(0)
	return res
}
