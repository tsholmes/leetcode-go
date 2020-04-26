package main

import "fmt"

func main() {
	fmt.Println(reformat("covid2019"))
	fmt.Println(reformat("leetcode"))
	fmt.Println(reformat("1234"))
	fmt.Println(reformat("1234abc"))
	fmt.Println(reformat("abcd1234"))
}
func reformat(s string) string {
	var ls []byte
	var ns []byte

	for _, c := range []byte(s) {
		if c >= 'a' && c <= 'z' {
			ls = append(ls, c)
		} else {
			ns = append(ns, c)
		}
	}

	if len(ls) > len(ns) {
		ns, ls = ls, ns
	}

	if len(ns) != len(ls)+1 && len(ns) != len(ls) {
		return ""
	}

	var res []byte
	for i := range ls {
		res = append(res, ns[i], ls[i])
	}
	if len(ns) > len(ls) {
		res = append(res, ns[len(ns)-1])
	}

	return string(res)
}
