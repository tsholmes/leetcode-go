package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"))
	fmt.Println(isValidSerialization("1,#"))
	fmt.Println(isValidSerialization("9,#,#,1"))
}

func isValidSerialization(preorder string) bool {
	ps := strings.Split(preorder, ",")
	if ps[0] == "#" {
		return len(ps) == 1
	}
	i := 0
	var valid func() bool
	valid = func() bool {
		if i == len(ps) {
			return false
		}
		v := ps[i]
		i++
		if v == "#" {
			return true
		}
		return valid() && valid()
	}
	return valid() && i == len(ps)
}
