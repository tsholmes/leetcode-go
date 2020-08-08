package main

func main() {

}

func countTriplets(arr []int) int {
	cumXor := make([]int, len(arr)+1)
	curXor := 0
	for i, n := range arr {
		curXor ^= n
		cumXor[i+1] = curXor
	}

	res := 0

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			x1 := cumXor[i] ^ cumXor[j]
			for k := j; k < len(arr); k++ {
				x2 := cumXor[j] ^ cumXor[k+1]
				if x1 == x2 {
					res++
				}
			}
		}
	}

	return res
}
