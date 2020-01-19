package main

func main() {

}

func canConstruct(ransomNote string, magazine string) bool {
	counts := map[byte]int{}
	for i := 0; i < len(ransomNote); i++ {
		counts[ransomNote[i]]++
	}
	total := len(ransomNote)
	for i := 0; i < len(magazine); i++ {
		c := magazine[i]
		if counts[c] > 0 {
			counts[c]--
			total--
		}
	}
	return total == 0
}
