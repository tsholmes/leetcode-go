package main

func main() {

}

func reverseVowels(s string) string {
	res := []byte(s)
	i, j := 0, len(s)-1
	vowel := func(b byte) bool {
		return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' ||
			b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U'
	}
	for i < j {
		for i < j && !vowel(res[i]) {
			i++
		}
		for j > i && !vowel(res[j]) {
			j--
		}
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return string(res)
}
