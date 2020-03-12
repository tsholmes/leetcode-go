package main

import "fmt"

func main() {
	fmt.Println(removeBoxes([]int{1, 3, 2, 2, 2, 3, 4, 3, 1}))
}

func removeBoxes(boxes []int) int {
	if len(boxes) == 0 {
		return 0
	}

	var bs [][2]int
	boxes = append(boxes, -1)
	last := boxes[0]
	count := 0
	for _, b := range boxes {
		if b == last {
			count++
		} else {
			bs = append(bs, [2]int{last, count})
			last, count = b, 1
		}
	}

	memo := map[[3]int]int{}
	var search func(i, j, ex int) int
	search = func(i, j, ex int) int {
		if i == j {
			return (ex + bs[i][1]) * (ex + bs[i][1])
		}
		key := [3]int{i, j, ex}
		if v, ok := memo[key]; ok {
			return v
		}

		max := search(i+1, j, 0) + (ex+bs[i][1])*(ex+bs[i][1])
		for k := i + 1; k <= j; k++ {
			if bs[k][0] == bs[i][0] {
				ss := search(i+1, k-1, 0) + search(k, j, ex+bs[i][1])
				if ss > max {
					max = ss
				}
			}
		}

		memo[key] = max
		return max
	}

	return search(0, len(bs)-1, 0)
}
