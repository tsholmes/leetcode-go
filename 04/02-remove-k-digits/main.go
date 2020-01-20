package main

import "fmt"

func main() {
	fmt.Println(removeKdigits("1432219", 3))
	fmt.Println(removeKdigits("10200", 1))
	fmt.Println(removeKdigits("10", 2))
	fmt.Println(removeKdigits("112", 1))
}

func removeKdigits(num string, k int) string {
	if k == len(num) {
		return "0"
	}
	var idxs [10][]int
	for i := 0; i < len(num); i++ {
		d := num[i] - '0'
		idxs[d] = append(idxs[d], i)
	}

	rem := 0
	i := 0
	var res []byte
	for rem < k && i < len(num) {
		min := 10
		for d := 0; d < 10; d++ {
			for len(idxs[d]) > 0 && idxs[d][0] < i {
				idxs[d] = idxs[d][1:]
			}
			if len(idxs[d]) > 0 && idxs[d][0]-i+rem <= k {
				min = d
				break
			}
		}
		res = append(res, byte(min)+'0')
		rem += idxs[min][0] - i
		i = idxs[min][0] + 1
	}
	res = append(res, num[i:]...)
	res = res[:len(num)-k]
	for len(res) > 1 && res[0] == '0' {
		res = res[1:]
	}
	return string(res)
}
