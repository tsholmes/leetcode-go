package main

import "fmt"

func main() {
	fmt.Println(canWin("++++"))
	fmt.Println(canWin("+++++++++++++++++++++"))
}

func canWin(s string) bool {
	b := []byte(s)
	type key struct {
		w string
		y bool
	}
	memo := map[key]bool{}
	var search func(bool) bool
	search = func(you bool) bool {
		k := key{string(b), you}
		if v, ok := memo[k]; ok {
			return v
		}
		all := !you
		found := false
		for i := 0; i < len(b)-1; i++ {
			j := i + 1
			if b[i] != '+' || b[j] != '+' {
				continue
			}
			found = true
			b[i], b[j] = '-', '-'
			if you {
				if search(!you) {
					all = true
					b[i], b[j] = '+', '+'
					break
				}
			} else {
				all = all && search(!you)
			}
			b[i], b[j] = '+', '+'
		}
		if !found {
			all = !you
		}
		memo[k] = all
		return all
	}
	return search(true)
}
