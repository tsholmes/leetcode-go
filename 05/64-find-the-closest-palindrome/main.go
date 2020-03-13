package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(nearestPalindromic("99"))
}

func nearestPalindromic(n string) string {
	pn, in, dn := toPal(n), toPal(incr(n)), toPal(decr(n))
	if pn == n {
		return minDiff(n, dn, in)
	} else if gt(pn, n) {
		return minDiff(n, dn, pn)
	} else {
		return minDiff(n, pn, in)
	}
}

func toPal(n string) string {
	bs := []byte(n)
	for i := 0; i < len(bs)/2; i++ {
		bs[len(bs)-i-1] = bs[i]
	}
	return string(bs)
}

func incr(n string) string {
	bs := []byte(n)
	mid := len(bs) / 2
	if len(bs)&1 == 0 {
		mid--
	}
	for i := mid; i >= 0; i-- {
		if bs[i] < '9' {
			bs[i]++
			break
		}
		bs[i] = '0'
	}
	if bs[0] == '0' {
		bs = append([]byte{'1'}, bs...)
	}
	return string(bs)
}

func decr(n string) string {
	if len(n) == 2 && n[0] == '1' {
		return "9"
	}
	bs := []byte(n)

	if bs[0] == '1' && len(bs) > 1 {
		zs := true
		for i := 1; i < len(bs); i++ {
			zs = zs && bs[i] == '0'
		}
		if zs {
			for i := 1; i < len(bs); i++ {
				bs[i] = '9'
			}
			return string(bs[1:])
		}
	}

	mid := len(bs) / 2
	if len(bs)&1 == 0 {
		mid--
	}
	for i := mid; i >= 0; i-- {
		if bs[i] > '0' {
			bs[i]--
			break
		}
		bs[i] = '9'
	}
	for len(bs) > 1 && bs[0] == '0' {
		bs = bs[1:]
	}
	return string(bs)
}

func minDiff(n, a, b string) string {
	if gt(a, b) {
		a, b = b, a
	}
	nn, _ := strconv.Atoi(n)
	an, _ := strconv.Atoi(a)
	bn, _ := strconv.Atoi(b)

	if abs(nn-an) <= abs(nn-bn) {
		return a
	}
	return b
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func gt(a, b string) bool {
	if len(a) > len(b) {
		return true
	} else if len(a) < len(b) {
		return false
	}
	return a > b
}
