package main

func main() {

}

func findTheDifference(s string, t string) byte {
	var counts [256]int
	for i := 0; i < len(s); i++ {
		counts[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		if counts[t[i]] == 0 {
			return t[i]
		}
		counts[t[i]]--
	}
	return 0
}
