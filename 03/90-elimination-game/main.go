package main

func main() {
	for i := 1; i <= 1<<16; i++ {
		r := lastRemaining(i)
		// fmt.Printf("%d %d\n", i, r)
		_ = r
	}
}

func lastRemaining(n int) int {
	base := 1
	add := 1
	for add<<1 <= n {
		base += add
		add <<= 1
		if n&add != 0 && add<<1 < n {
			base += add
		}
		add <<= 1
	}

	return base
}
