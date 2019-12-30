package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(generatePalindromes("aabb"))
	fmt.Println(generatePalindromes("abc"))
	fmt.Println(generatePalindromes("aaaaaa"))
	fmt.Println(generatePalindromes("aaaabb"))
}

func generatePalindromes(s string) []string {
	bs := []byte(s)
	sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })

	var mid byte
	count := 0
	for i := 0; i < len(bs); i++ {
		if i == len(bs)-1 {
			if mid != 0 {
				return nil
			}
			mid = bs[i]
			break
		}
		if bs[i] != bs[i+1] {
			if mid != 0 {
				return nil
			}
			mid = bs[i]
			continue
		}
		bs[count] = bs[i]
		count++
		i++
	}
	bs = bs[:count]

	if mid != 0 && len(s)&1 == 0 {
		return nil
	}

	var res []string
	work := make([]byte, len(s))
	if len(s)&1 != 0 {
		work[len(s)/2] = mid
	}

	var search func(int, int)
	search = func(i int, used int) {
		j := len(s) - i - 1
		if j <= i {
			res = append(res, string(work))
			return
		}

		var last byte
		for k := range bs {
			b := 1 << uint(k)
			if used&b != 0 {
				continue
			}
			if bs[k] == last {
				continue
			}
			last = bs[k]
			work[i], work[j] = bs[k], bs[k]
			search(i+1, used|b)
		}
	}

	search(0, 0)

	return res
}
