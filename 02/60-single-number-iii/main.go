package main

func main() {

}

func singleNumber(nums []int) []int {
	xor := 0
	for _, n := range nums {
		xor ^= n
	}
	low := xor & -xor
	a := 0
	for _, n := range nums {
		if n&low != 0 {
			a ^= n
		}
	}
	return []int{a, xor ^ a}
}
