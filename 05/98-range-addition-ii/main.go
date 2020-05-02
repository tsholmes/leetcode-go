package main

func main() {

}

func maxCount(m int, n int, ops [][]int) int {
	mina, minb := m, n
	for _, op := range ops {
		if op[0] < mina {
			mina = op[0]
		}
		if op[1] < minb {
			minb = op[1]
		}
	}
	return mina * minb
}
