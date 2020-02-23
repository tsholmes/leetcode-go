package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(minTransfers([][]int{
		{0, 1, 10}, {2, 0, 5},
	}))
	fmt.Println(minTransfers([][]int{
		{0, 1, 10}, {1, 0, 1}, {1, 2, 5}, {2, 0, 5},
	}))
}

func minTransfers(transactions [][]int) int {
	idMap := map[int]int{}
	addID := func(id int) int {
		if v, ok := idMap[id]; ok {
			return v
		}
		v := len(idMap)
		idMap[id] = v
		return v
	}
	for _, t := range transactions {
		addID(t[0])
		addID(t[1])
	}
	N := len(idMap)

	bals := make([]int, N)
	for _, t := range transactions {
		i, j, v := idMap[t[0]], idMap[t[1]], t[2]
		bals[i] -= v
		bals[j] += v
	}
	bals = reduce(bals)

	count := 0
	for len(bals) > 0 {
		minMask := (1 << uint(len(bals))) - 1
		minC := len(bals)
		for mask := 1; mask < 1<<uint(len(bals)); mask++ {
			bc := bits.OnesCount(uint(mask))
			if bc >= minC {
				continue
			}
			sum := 0
			for i, n := range bals {
				if mask&(1<<uint(i)) != 0 {
					sum += n
				}
			}
			if sum == 0 {
				minMask, minC = mask, bc
			}
		}
		count += minC - 1
		rcount := 0
		for i, n := range bals {
			if minMask&(1<<uint(i)) == 0 {
				bals[rcount] = n
				rcount++
			}
		}
		bals = bals[:rcount]
	}

	return count
}

func reduce(bals []int) []int {
	count := 0
	for _, n := range bals {
		if n != 0 {
			bals[count] = n
			count++
		}
	}
	return bals[:count]
}
