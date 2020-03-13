package main

func main() {

}

func leastBricks(wall [][]int) int {
	edges := map[int]int{}
	for _, row := range wall {
		wid := 0
		for _, b := range row[:len(row)-1] {
			wid += b
			edges[wid]++
		}
	}
	max := 0
	for _, v := range edges {
		if v > max {
			max = v
		}
	}
	return len(wall) - max
}
