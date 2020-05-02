package main

import "fmt"

func main() {
	fmt.Println(maxScore([]int{96, 90, 41, 82, 39, 74, 64, 50, 30}, 8))
}

func maxScore(cardPoints []int, k int) int {
	N := len(cardPoints)
	start := make([]int, N+1)
	end := make([]int, N+1)
	for i := 0; i < N; i++ {
		start[i+1] = start[i] + cardPoints[i]
		end[i+1] = end[i] + cardPoints[N-i-1]
	}

	max := 0
	for i := 0; i <= k; i++ {
		s := start[i] + end[k-i]
		if s > max {
			max = s
		}
	}

	return max
}
