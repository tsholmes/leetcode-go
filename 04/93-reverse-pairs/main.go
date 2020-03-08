package main

import (
	"fmt"
	"math"
)

func main() {
	// ns := make([]int, 50000)
	// for i := range ns {
	// 	ns[i] = 50000 - i
	// }
	// start := time.Now()
	// reversePairs(ns)
	// end := time.Now()
	// fmt.Println(end.Sub(start))
	fmt.Println(reversePairs([]int{1, 3, 2, 3, 1}))
	fmt.Println(reversePairs([]int{2, 4, 3, 5, 1}))
}

func reversePairs(nums []int) int {
	root := &bitNode{}

	norm := func(n int) int {
		return n - (2 * math.MinInt32)
	}

	count := 0
	for _, n := range nums {
		count += root.countGreater(norm(n*2), 1<<34)
		root.add(norm(n), 1<<34)
	}

	return count
}

type bitNode struct {
	count int
	cs    [2]*bitNode
}

func (b *bitNode) add(n int, bit int) {
	b.count++
	if bit == 0 {
		return
	}
	if n&bit == 0 {
		if b.cs[0] == nil {
			b.cs[0] = &bitNode{}
		}
		b.cs[0].add(n, bit>>1)
	} else {
		if b.cs[1] == nil {
			b.cs[1] = &bitNode{}
		}
		b.cs[1].add(n, bit>>1)
	}
}

func (b *bitNode) countGreater(n int, bit int) int {
	if b == nil {
		return 0
	}
	if bit == 0 {
		return 0
	}
	if n&bit != 0 {
		return b.cs[1].countGreater(n, bit>>1)
	} else {
		res := 0
		if b.cs[1] != nil {
			res += b.cs[1].count
		}
		res += b.cs[0].countGreater(n, bit>>1)
		return res
	}
}
