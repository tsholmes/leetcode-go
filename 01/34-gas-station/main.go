package main

import "fmt"

func main() {
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
}

func canCompleteCircuit(gas []int, cost []int) int {
	for i := range gas {
		valid := true
		curGas := 0
		for p := 0; p < len(gas); p++ {
			j := (i + p) % len(gas)
			curGas += gas[j]
			curGas -= cost[j]
			if curGas < 0 {
				valid = false
				break
			}
		}
		if valid {
			return i
		}
	}
	return -1
}
