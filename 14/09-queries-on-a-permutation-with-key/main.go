package main

func main() {

}

func processQueries(queries []int, m int) []int {
	nums := make([]int, m)
	for i := range nums {
		nums[i] = i + 1
	}

	var res []int
	var next []int
	for _, q := range queries {
		next = next[:0]
		idx := -1
		next = append(next, q)
		for i, n := range nums {
			if n == q {
				idx = i
			} else {
				next = append(next, n)
			}
		}
		res = append(res, idx)
		nums, next = next, nums
	}
	return res
}
