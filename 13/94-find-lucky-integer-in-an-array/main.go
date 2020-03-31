package main

func main() {

}
func findLucky(arr []int) int {
	max := -1
	counts := map[int]int{}
	for _, n := range arr {
		counts[n]++
	}

	for k, v := range counts {
		if k == v && k > max {
			max = k
		}
	}
	return max
}
