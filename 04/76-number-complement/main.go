package main

import "math/bits"

func main() {

}

func findComplement(num int) int {
	if num == 0 {
		return 1
	}
	mask := (1 << uint(63-bits.LeadingZeros(uint(num)))) - 1
	return mask &^ num
}
