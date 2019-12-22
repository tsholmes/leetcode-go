package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(numbers []int, target int) []int {
	i1, i2 := 0, len(numbers)-1
	for i1 < i2 {
		s := numbers[i1] + numbers[i2]
		if s > target {
			i2--
		} else if s < target {
			i1++
		} else {
			return []int{i1 + 1, i2 + 1}
		}
	}
	return nil
}
