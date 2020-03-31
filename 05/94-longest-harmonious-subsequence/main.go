package main

func main() {

}

func findLHS(nums []int) int {
	counts := map[int]int{}
	for _, n := range nums {
		counts[n]++
	}
	max := 0
	for k, v := range counts {
		if vl := v + counts[k-1]; vl > v && vl > max {
			max = vl
		}
		if vh := v + counts[k+1]; vh > v && vh > max {
			max = vh
		}
	}
	return max
}
