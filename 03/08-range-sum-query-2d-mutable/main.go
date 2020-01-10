package main

func main() {

}

type NumMatrix struct {
	rows []NumArray
}

func Constructor(matrix [][]int) NumMatrix {
	rows := make([]NumArray, len(matrix))
	for i, row := range matrix {
		rows[i] = RConstructor(row)
	}
	return NumMatrix{rows: rows}
}

func (n *NumMatrix) Update(row int, col int, val int) {
	n.rows[row].Update(col, val)
}

func (n *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i <= row2; i++ {
		sum += n.rows[i].SumRange(col1, col2)
	}
	return sum
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * obj.Update(row,col,val);
 * param_2 := obj.SumRegion(row1,col1,row2,col2);
 */

type rangeSum struct {
	sum       int
	i, j, mid int
	left      int
	right     int
}

type NumArray struct {
	sums []rangeSum
}

func RConstructor(nums []int) NumArray {
	var sums []rangeSum
	var build func(i, j int) int
	build = func(i, j int) int {
		if j < i {
			return 0
		}
		sum := 0
		left, right := -1, -1
		mid := i
		if i == j {
			sum = nums[i]
		} else {
			mid = (i + j) / 2
			left = build(i, mid)
			right = build(mid+1, j)
			sum = sums[left].sum + sums[right].sum
		}
		idx := len(sums)
		sums = append(sums, rangeSum{
			sum:   sum,
			i:     i,
			j:     j,
			mid:   mid,
			left:  left,
			right: right,
		})
		return idx
	}
	build(0, len(nums)-1)
	return NumArray{
		sums: sums,
	}
}

func (n *NumArray) Update(i int, val int) {
	idx := len(n.sums) - 1
	prevVal := 0
	for {
		r := n.sums[idx]
		if r.i == r.j {
			prevVal = r.sum
			break
		}
		if i <= r.mid {
			idx = r.left
		} else {
			idx = r.right
		}
	}
	idx = len(n.sums) - 1
	for {
		r := &n.sums[idx]
		r.sum += val - prevVal
		if r.i == r.j {
			break
		}
		if i <= r.mid {
			idx = r.left
		} else {
			idx = r.right
		}
	}
}

func (n *NumArray) sumUpTo(i int) int {
	idx := len(n.sums) - 1
	sum := 0
	for {
		r := n.sums[idx]
		if r.j == i {
			sum += r.sum
			break
		}
		if i <= r.mid {
			idx = r.left
			continue
		}
		sum += n.sums[r.left].sum
		idx = r.right
	}
	return sum
}

func (n *NumArray) SumRange(i int, j int) int {
	sum := n.sumUpTo(j)
	if i > 0 {
		sum -= n.sumUpTo(i - 1)
	}
	return sum
}
