package main

import "fmt"

func main() {
	fmt.Println(strobogrammaticInRange("50", "100"))
	fmt.Println(strobogrammaticInRange("100", "50"))
	fmt.Println(strobogrammaticInRange("100", "100"))
	fmt.Println(strobogrammaticInRange("0", "99999"))
	fmt.Println(strobogrammaticInRange("0", "0"))
}

func strobogrammaticInRange(low string, high string) int {
	if len(low) > len(high) || (len(low) == len(high) && low > high) {
		return 0
	}
	countTo := make([]int, len(high)+1)
	countTo[1] = 3
	if len(high) > 2 {
		countTo[2] = 4
		if len(high) > 3 {
			countTo[3] = countTo[1] * 4
		}
	}
	for i := 4; i <= len(high); i++ {
		countTo[i] = countTo[i-2] * 5
	}
	for i := 1; i <= len(high); i++ {
		countTo[i] += countTo[i-1]
	}

	same := []byte(`018`)
	var countUp func(string, bool, int, int, []byte) int
	countUp = func(s string, lt bool, i int, count int, work []byte) int {
		j := len(s) - i - 1
		if j < i {
			if !lt {
				if string(work) <= s {
					return count
				}
				return 0
			}
			return count
		}
		if j == i {
			if lt {
				return count * 3
			}
			mul := 0
			for _, c := range same {
				work[i] = c
				if string(work) <= s {
					mul++
				}
			}
			return mul * count
		}

		if lt {
			return countUp(s, lt, i+1, count*5, work)
		}

		zcount := 0
		if i > 0 {
			zcount++
		}

		switch s[i] {
		case '9':
			work[i] = '9'
			work[j] = '6'
			return countUp(s, false, i+1, count, work) + countUp(s, true, i+1, count*(3+zcount), work)
		case '8':
			work[i] = '8'
			work[j] = '8'
			return countUp(s, false, i+1, count, work) + countUp(s, true, i+1, count*(2+zcount), work)
		case '7':
			return countUp(s, true, i+1, count*(2+zcount), work)
		case '6':
			work[i] = '6'
			work[j] = '9'
			return countUp(s, false, i+1, count, work) + countUp(s, true, i+1, count*(1+zcount), work)
		case '5':
			return countUp(s, true, i+1, count*(1+zcount), work)
		case '4':
			return countUp(s, true, i+1, count*(1+zcount), work)
		case '3':
			return countUp(s, true, i+1, count*(1+zcount), work)
		case '2':
			return countUp(s, true, i+1, count*(1+zcount), work)
		case '1':
			work[i] = '1'
			work[j] = '1'
			res := countUp(s, false, i+1, count, work)
			if i > 0 {
				res += countUp(s, true, i+1, count, work)
			}
			return res
		case '0':
			work[i] = '0'
			work[j] = '0'
			return countUp(s, false, i+1, count, work)
		}

		return 0
	}

	if low == "0" {
		return countUp(high, false, 0, 1, make([]byte, len(high))) + countTo[len(high)-1]
	}

	lowb := []byte(low)
	for i := len(lowb) - 1; i >= 0; i-- {
		if lowb[i] > '0' {
			lowb[i]--
			break
		}
		lowb[i] = '9'
	}
	for len(lowb) > 1 && lowb[0] == '0' {
		lowb = lowb[1:]
	}
	low = string(lowb)

	return countUp(high, false, 0, 1, make([]byte, len(high))) - countUp(low, false, 0, 1, make([]byte, len(low))) + countTo[len(high)-1] - countTo[len(low)-1]
}
