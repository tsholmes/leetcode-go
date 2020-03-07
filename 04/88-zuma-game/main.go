package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(insert("WWRRWW", 'R', 2))
}

func findMinStep(board string, hand string) int {
	memo := map[[2]string]int{}
	var search func(b, h string) int
	search = func(b, h string) int {
		if len(b) == 0 {
			return 0
		}
		if len(h) == 0 {
			return math.MaxInt32
		}

		k := [2]string{b, h}
		if v, ok := memo[k]; ok {
			return v
		}

		min := math.MaxInt32
		for i := 0; i < len(h); i++ {
			for j := 0; j <= len(b); j++ {
				v := 1 + search(insert(b, h[i], j), remove(h, i))
				if v < min {
					min = v
				}
			}
		}
		memo[k] = min
		return min
	}

	res := search(board, hand)
	if res >= math.MaxInt32 {
		return -1
	}
	return res
}

func insert(board string, b byte, i int) string {
	var res []byte
	res = append(res, board[:i]...)
	res = append(res, b)
	res = append(res, board[i:]...)

	for {
		count := 0
		for i := 0; i < len(res); i++ {
			if i+2 < len(res) && res[i] == res[i+1] && res[i] == res[i+2] {
				j := i
				for j < len(res) && res[i] == res[j] {
					j++
				}
				i = j - 1
			} else {
				res[count] = res[i]
				count++
			}
		}
		if count == len(res) {
			break
		}
		res = res[:count]
	}
	return string(res)
}

func remove(hand string, i int) string {
	return string(append([]byte(hand[:i]), hand[i+1:]...))
}
