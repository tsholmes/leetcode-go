package main

func main() {

}

type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{}
	}
	N, M := len(matrix), len(matrix[0])

	sums := make([][]int, N)
	for i := range matrix {
		sums[i] = make([]int, M)
		for j := range matrix[i] {
			sums[i][j] = matrix[i][j]
			switch {
			case i > 0 && j > 0:
				sums[i][j] += sums[i-1][j]
				sums[i][j] += sums[i][j-1]
				sums[i][j] -= sums[i-1][j-1]
			case i > 0:
				sums[i][j] += sums[i-1][j]
			case j > 0:
				sums[i][j] += sums[i][j-1]
			}
		}
	}
	return NumMatrix{sums: sums}
}

func (n *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := n.sums[row2][col2]
	switch {
	case row1 > 0 && col1 > 0:
		sum -= n.sums[row1-1][col2]
		sum -= n.sums[row2][col1-1]
		sum += n.sums[row1-1][col1-1]
	case row1 > 0:
		sum -= n.sums[row1-1][col2]
	case col1 > 0:
		sum -= n.sums[row2][col1-1]
	}
	return sum
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
