package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(maximum69Number(9996))
	fmt.Println(maximum69Number(9999))
}

func maximum69Number(num int) int {
	bs := []byte(strconv.FormatInt(int64(num), 10))
	for i := range bs {
		if bs[i] == '6' {
			bs[i] = '9'
			break
		}
	}
	v, _ := strconv.ParseInt(string(bs), 10, 64)
	return int(v)
}
