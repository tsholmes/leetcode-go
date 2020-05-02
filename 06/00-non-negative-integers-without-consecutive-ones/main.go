package main

func main() {

}

func findIntegers(num int) int {
	type key struct {
		bit   int
		last1 bool
		max   bool
	}
	memo := map[key]int{}
	var search func(bit int, last1, max bool) int
	search = func(bit int, last1, max bool) int {
		if bit == -1 {
			return 1
		}

		k := key{bit, last1, max}
		if v, ok := memo[k]; ok {
			return v
		}

		var res int

		if !last1 && (!max || (num&(1<<uint(bit))) != 0) {
			res += search(bit-1, true, max)
		}
		res += search(bit-1, false, max && (num&(1<<uint(bit))) == 0)

		memo[k] = res
		return res
	}

	return search(31, false, true)
}
