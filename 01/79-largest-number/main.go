package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(largestNumber([]int{10, 2}))
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
	fmt.Println(largestNumber([]int{12, 1}))
	fmt.Println(largestNumber([]int{0, 0, 0}))
}

func largestNumber(nums []int) string {
	ss := make([]string, len(nums))
	for i, n := range nums {
		ss[i] = strconv.Itoa(n)
	}
	sort.Slice(ss, func(i, j int) bool {
		return (ss[i] + ss[j]) > (ss[j] + ss[i])
	})
	if len(ss) == 0 || ss[0] == "0" {
		return "0"
	}
	return strings.Join(ss, "")
}
