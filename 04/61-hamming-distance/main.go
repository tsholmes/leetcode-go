package main

import "math/bits"

func main() {

}

func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}
