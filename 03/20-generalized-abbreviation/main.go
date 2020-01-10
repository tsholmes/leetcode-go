package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(generateAbbreviations("word"))
}

func generateAbbreviations(word string) []string {
	var res []string
	var work []byte
	var search func(int, bool)
	search = func(i int, skip bool) {
		if i == len(word) {
			res = append(res, string(work))
			return
		}
		wlen := len(work)

		work = append(work, word[i])
		search(i+1, false)
		work = work[:wlen]

		if !skip {
			for n := 1; i+n <= len(word); n++ {
				work = strconv.AppendInt(work, int64(n), 10)
				search(i+n, true)
				work = work[:wlen]
			}
		}
	}
	search(0, false)
	return res
}
