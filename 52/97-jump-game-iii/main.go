package main

import "fmt"

func main() {
	fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 5))
	fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 0))
	fmt.Println(canReach([]int{3, 0, 2, 1, 2}, 2))
}

func canReach(arr []int, start int) bool {
	visited := make([]bool, len(arr))

	var search func(int) bool
	search = func(i int) bool {
		if i < 0 || i >= len(arr) {
			return false
		}
		if visited[i] {
			return false
		}
		visited[i] = true
		if arr[i] == 0 {
			return true
		}
		if search(i-arr[i]) || search(i+arr[i]) {
			return true
		}
		return false
	}

	return search(start)
}
