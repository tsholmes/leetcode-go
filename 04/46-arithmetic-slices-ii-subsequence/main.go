package main

func main() {

}

func numberOfArithmeticSlices(A []int) int {
	counts := map[[2]int]int{}
	for i := 1; i < len(A); i++ {
		for j := 0; j < i; j++ {
			counts[[2]int{i, A[i] - A[j]}] += counts[[2]int{j, A[i] - A[j]}] + 1
		}
	}
	res := 0
	for _, v := range counts {
		res += v
	}
	n := len(A)
	return res - (n * (n - 1) / 2)
}
