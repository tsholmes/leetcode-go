package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

}

func complexNumberMultiply(a string, b string) string {
	parse := func(s string) [2]int {
		mid := strings.IndexByte(s, '+')
		r, _ := strconv.Atoi(s[:mid])
		i, _ := strconv.Atoi(s[mid+1 : len(s)-1])
		return [2]int{r, i}
	}
	ai, bi := parse(a), parse(b)
	res := [2]int{ai[0]*bi[0] - ai[1]*bi[1], ai[0]*bi[1] + ai[1]*bi[0]}
	return fmt.Sprintf("%d+%di", res[0], res[1])
}
