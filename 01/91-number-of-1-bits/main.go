package main

func main() {

}

func hammingWeight(num uint32) int {
	num = (num & 0x55555555) + ((num >> 1) & 0x55555555)
	num = (num & 0x33333333) + ((num >> 2) & 0x33333333)
	num = (num & 0x0F0F0F0F) + ((num >> 4) & 0x0F0F0F0F)
	num = (num & 0x00FF00FF) + ((num >> 8) & 0x00FF00FF)
	num = (num & 0x0000FFFF) + ((num >> 16) & 0x0000FFFF)
	return int(num)
}
