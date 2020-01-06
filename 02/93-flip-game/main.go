package main

func main() {

}

func generatePossibleNextMoves(s string) []string {
	var res []string
	b := []byte(s)
	for i := 0; i < len(b)-1; i++ {
		j := i + 1
		if b[i] != '+' || b[j] != '+' {
			continue
		}
		b[i], b[j] = '-', '-'
		res = append(res, string(b))
		b[i], b[j] = '+', '+'
	}
	return res
}
