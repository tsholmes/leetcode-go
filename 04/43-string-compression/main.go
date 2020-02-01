package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(compress([]byte(`aabbccc`)))
}

func compress(chars []byte) int {
	if len(chars) < 2 {
		return len(chars)
	}
	count := 0
	lastc := chars[0]
	ccount := 1
	for i := 1; i < len(chars); i++ {
		if chars[i] != chars[i-1] {
			chars[count] = lastc
			count++
			if ccount > 1 {
				count += len(strconv.AppendInt(chars[count:count], int64(ccount), 10))
			}
			lastc = chars[i]
			ccount = 1
		} else {
			ccount++
		}
	}
	chars[count] = lastc
	count++
	if ccount > 1 {
		count += len(strconv.AppendInt(chars[count:count], int64(ccount), 10))
	}

	return count
}
