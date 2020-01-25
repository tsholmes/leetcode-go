package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(strongPasswordChecker(""))
	fmt.Println(strongPasswordChecker("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"))
}

func strongPasswordChecker(s string) int {
	type key struct {
		idx   int
		len   int
		lower bool
		upper bool
		digit bool
		last  byte
		run   int
	}
	cs := []byte(`012ABCabc`)
	min := math.MaxInt64
	memo := map[key]int{}

	update := func(k key, c byte, incr bool) key {
		if c == k.last {
			k.run++
		} else {
			k.last = c
			k.run = 1
		}
		if c >= '0' && c <= '9' {
			k.digit = true
		} else if c >= 'A' && c <= 'Z' {
			k.upper = true
		} else if c >= 'a' && c <= 'z' {
			k.lower = true
		}
		k.len++
		if incr {
			k.idx++
		}
		return k
	}

	var search func(key, int)
	search = func(k key, cost int) {
		if k.len > 20 || k.run > 2 {
			return
		}
		if cost >= min {
			return
		}
		if v, ok := memo[k]; ok && v <= cost {
			return
		}
		memo[k] = cost
		if k.idx == len(s) {
			missing := 0
			if !k.lower {
				missing++
			}
			if !k.upper {
				missing++
			}
			if !k.digit {
				missing++
			}
			if k.len < 6 {
				cost += 6 - k.len
				missing -= 6 - k.len
			}
			if k.len <= 20 && missing <= 0 && k.run < 3 && cost < min {
				min = cost
			}
			return
		}

		// keep
		if k.len < 20 {
			search(update(k, s[k.idx], true), cost)
		}

		// replace
		if k.len < 20 {
			for _, c := range cs {
				if c == s[k.idx] {
					continue
				}
				search(update(k, c, true), cost+1)
			}
		}

		// remove
		{
			k2 := k
			k2.idx++
			search(k2, cost+1)
		}

		// insert
		if k.len < 20 {
			for _, c := range cs {
				search(update(k, c, false), cost+1)
			}
		}
	}
	search(key{}, 0)
	return min
}
