package main

func main() {

}

func splitArray(nums []int) bool {
	N := len(nums)
	if N < 7 {
		return false
	}
	totalSum := 0
	sums := make([]int, len(nums))
	for i, n := range nums {
		totalSum += n
		sums[i] = totalSum
	}

	rangeSum := func(i, j int) int {
		s := sums[j]
		if i > 0 {
			s -= sums[i-1]
		}
		return s
	}

	for mid := 3; mid < N-3; mid++ {
		lsums := map[int]bool{}
		for l := 1; l < mid-1; l++ {
			l1, l2 := rangeSum(0, l-1), rangeSum(l+1, mid-1)
			if l1 == l2 {
				lsums[l1] = true
			}
		}
		for r := mid + 2; r < N-1; r++ {
			r1, r2 := rangeSum(mid+1, r-1), rangeSum(r+1, N-1)
			if r1 == r2 && lsums[r1] {
				return true
			}
		}
	}

	return false
}
