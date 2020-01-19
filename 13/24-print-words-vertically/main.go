package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(printVertically("HOW ARE YOU"))
	fmt.Println(printVertically("CONTEST IS COMING"))
}

func printVertically(s string) []string {
	if s == "" {
		return nil
	}
	ws := strings.Split(s, " ")
	max := 0
	for _, w := range ws {
		if len(w) > max {
			max = len(w)
		}
	}
	res := make([][]byte, max)
	for i := range res {
		res[i] = make([]byte, len(ws))
		for j := range res[i] {
			res[i][j] = ' '
		}
	}

	for i, w := range ws {
		for j := 0; j < len(w); j++ {
			res[j][i] = w[j]
		}
	}

	ss := make([]string, len(res))
	for i, r := range res {
		for r[len(r)-1] == ' ' {
			r = r[:len(r)-1]
		}
		ss[i] = string(r)
	}
	return ss
}
