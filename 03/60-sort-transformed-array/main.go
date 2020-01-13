package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sortTransformedArray(
		[]int{-99, -94, -90, -88, -84, -83, -79, -68, -58, -52, -52, -50, -47, -45, -35, -29, -5, -1, 9, 12, 13, 25, 27, 33, 36, 38, 40, 46, 47, 49, 57, 57, 61, 63, 73, 75, 79, 97, 98},
		-2,
		44,
		-56,
	))
}

func sortTransformedArray(nums []int, a int, b int, c int) []int {
	if len(nums) == 0 {
		return nil
	}
	var mid float64
	if a == 0 {
		mid = -math.MaxFloat64
	} else {
		mid = float64(-b) / float64(2*a)
	}
	var lt []int
	var gt []int
	for _, n := range nums {
		v := a*n*n + b*n + c
		if float64(n) < mid {
			lt = append(lt, v)
		} else {
			gt = append(gt, v)
		}
	}

	if len(lt) > 0 && lt[0] > lt[len(lt)-1] {
		for i := 0; i < len(lt)/2; i++ {
			j := len(lt) - i - 1
			lt[i], lt[j] = lt[j], lt[i]
		}
	}

	if len(gt) > 0 && gt[0] > gt[len(gt)-1] {
		for i := 0; i < len(gt)/2; i++ {
			j := len(gt) - i - 1
			gt[i], gt[j] = gt[j], gt[i]
		}
	}

	i, j := 0, 0
	var res []int
	for i < len(lt) && j < len(gt) {
		vi, vj := lt[i], gt[j]
		if vi < vj {
			res = append(res, vi)
			i++
		} else {
			res = append(res, vj)
			j++
		}
	}
	res = append(res, lt[i:]...)
	res = append(res, gt[j:]...)
	return res
}
