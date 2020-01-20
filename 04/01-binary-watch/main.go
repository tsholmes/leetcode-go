package main

import (
	"fmt"
	"math/bits"
)

func main() {

}

func readBinaryWatch(num int) []string {
	var res []string
	for h := 0; h < 12; h++ {
		for m := 0; m < 60; m++ {
			if bits.OnesCount8(uint8(h))+bits.OnesCount8(uint8(m)) == num {
				res = append(res, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return res
}
