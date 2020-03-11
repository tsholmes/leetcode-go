package main

import (
	"math/bits"
	"strings"
)

func main() {

}

func findWords(words []string) []string {
	rows := []string{
		"QWERTYUIOP",
		"ASDFGHJKL",
		"ZXCVBNM",
	}

	crows := map[byte]int{}
	for i, r := range rows {
		for j := 0; j < len(r); j++ {
			crows[r[j]] = i
		}
	}

	count := 0
	for _, w := range words {
		wu := strings.ToUpper(w)
		mask := 0
		for i := 0; i < len(wu); i++ {
			mask |= 1 << uint(crows[wu[i]])
		}
		if bits.OnesCount(uint(mask)) == 1 {
			words[count] = w
			count++
		}
	}

	return words[:count]
}
