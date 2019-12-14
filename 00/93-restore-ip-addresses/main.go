package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses(""))
	fmt.Println(restoreIpAddresses("0000"))
	fmt.Println(restoreIpAddresses("010010"))
}

func restoreIpAddresses(s string) []string {
	var res []string
	work := make([]string, 4)
	var search func(int, int)
	search = func(i, j int) {
		if j == 4 {
			if i == len(s) {
				res = append(res, strings.Join(work, "."))
			}
			return
		} else if i == len(s) {
			return
		}

		for k := i + 1; k <= i+3 && k <= len(s); k++ {
			v, _ := strconv.ParseInt(s[i:k], 10, 64)
			if k-i > 1 && s[i] == '0' {
				continue
			}
			if v >= 0 && v <= 255 {
				work[j] = s[i:k]
				search(k, j+1)
			}
		}
	}

	search(0, 0)
	return res
}
