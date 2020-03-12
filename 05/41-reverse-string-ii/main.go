package main

func main() {

}

func reverseStr(s string, k int) string {
	revk := func(b []byte) {
		if len(b) > k {
			b = b[:k]
		}
		for i := 0; i < len(b)/2; i++ {
			j := len(b) - i - 1
			b[i], b[j] = b[j], b[i]
		}
	}
	bs := []byte(s)
	for i := 0; i < len(s); i += k + k {
		revk(bs[i:])
	}
	return string(bs)
}
