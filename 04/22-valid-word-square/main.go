package main

func main() {

}

func validWordSquare(words []string) bool {
	for i := range words {
		for j := 0; j < len(words[i]); j++ {
			if j >= len(words) {
				return false
			}
			if i >= len(words[j]) {
				return false
			}
			if words[i][j] != words[j][i] {
				return false
			}
		}
	}
	return true
}
