package main

func main() {

}

func matrixReshape(nums [][]int, r int, c int) [][]int {
	R, C := len(nums), len(nums[0])
	if R*C != r*c {
		return nums
	}
	ii, ij := 0, 0
	res := make([][]int, r)
	for i := range res {
		res[i] = make([]int, c)
		for j := range res[i] {
			res[i][j] = nums[ii][ij]
			ij++
			if ij == C {
				ii++
				ij = 0
			}
		}
	}
	return res
}
