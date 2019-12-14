package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 30; i++ {
		fmt.Println(countAndSay(i))
	}
}

func countAndSay(n int) string {
	current := make([]byte, 0, n)
	current = append(current, '1')
	next := make([]byte, 0, n)

	for i := 1; i < n; i++ {
		last := current[0]
		lastCount := 1
		for j := 1; j < len(current); j++ {
			ch := current[j]
			if ch == last {
				lastCount++
			} else {
				next = strconv.AppendInt(next, int64(lastCount), 10)
				next = append(next, last)
				last = ch
				lastCount = 1
			}
		}
		next = strconv.AppendInt(next, int64(lastCount), 10)
		next = append(next, last)

		current, next = next, current[:0]
	}

	return string(current)
}
