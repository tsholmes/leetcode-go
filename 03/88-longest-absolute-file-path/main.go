package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthLongestPath("dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"))
}

func lengthLongestPath(input string) int {
	lines := strings.Split(input, "\n")
	max := 0
	depth := 0
	sums := []int{-1}
	for _, l := range lines {
		d := 1
		for l[0] == '\t' {
			l = l[1:]
			d++
		}
		if d > depth {
			sums = append(sums, sums[depth]+1+len(l))
		} else {
			sums[d] = sums[d-1] + 1 + len(l)
			sums = sums[:d+1]
		}
		depth = d
		if idx := strings.IndexByte(l, '.'); idx != -1 && idx < len(l)-1 {
			if sums[d] > max {
				max = sums[d]
			}
		}
	}
	return max
}
