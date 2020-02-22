package main

func main() {

}

func countNegatives(grid [][]int) int {
	count := 0
	for _, row := range grid {
		for _, v := range row {
			if v < 0 {
				count++
			}
		}
	}
	return count
}
