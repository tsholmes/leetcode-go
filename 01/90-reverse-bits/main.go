package main

func main() {
}

func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		bi := uint32(1) << uint(i)
		bj := uint32(1) << uint(32-i-1)
		if num&bi != 0 {
			res |= bj
		}
	}
	return res
}
