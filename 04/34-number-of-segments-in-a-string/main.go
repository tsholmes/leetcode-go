package main

func main() {

}

func countSegments(s string) int {
	var last byte = ' '
	count := 0
	for i := 0; i < len(s); i++ {
		if last == ' ' && s[i] != ' ' {
			count++
		}
		last = s[i]
	}
	return count
}
