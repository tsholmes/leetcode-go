package main

import "fmt"

func main() {
	fmt.Println(nthSuperUglyNumber(12, []int{2, 7, 13, 19}))
}

func nthSuperUglyNumber(n int, primes []int) int {
	nums := make([]int, 0, n)
	nums = append(nums, 1)
	ps := make([]int, len(primes))
	for i := 1; i < n; i++ {
		min := nums[ps[0]] * primes[0]
		for pi, p := range primes {
			v := nums[ps[pi]] * p
			if v < min {
				min = v
			}
		}
		nums = append(nums, min)
		for pi, p := range primes {
			if nums[ps[pi]]*p == min {
				ps[pi]++
			}
		}
	}
	return nums[n-1]
}
