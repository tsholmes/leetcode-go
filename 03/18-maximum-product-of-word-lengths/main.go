package main

func main() {

}

func maxProduct(words []string) int {
	masks := make([]int, len(words))
	for i, w := range words {
		m := 0
		for j := 0; j < len(w); j++ {
			m |= 1 << (w[j] - 'a')
		}
		masks[i] = m
	}
	max := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if masks[i]&masks[j] != 0 {
				continue
			}
			l := len(words[i]) * len(words[j])
			if l > max {
				max = l
			}
		}
	}
	return max
}
