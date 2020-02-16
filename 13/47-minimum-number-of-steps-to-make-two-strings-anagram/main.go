package main

func main() {

}

func minSteps(s string, t string) int {
	var c1s [26]int
	var c2s [26]int
	for i := 0; i < len(s); i++ {
		c1s[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		c2s[t[i]-'a']++
	}

	count := 0

	for i := 0; i < 26; i++ {
		d := c1s[i] - c2s[i]
		if d < 0 {
			d = -d
		}
		count += d
	}

	return count / 2
}
