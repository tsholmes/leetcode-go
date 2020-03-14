package main

func main() {

}

func checkInclusion(s1 string, s2 string) bool {
	var exp [26]int
	for i := 0; i < len(s1); i++ {
		exp[s1[i]-'a']++
	}

	seen := map[[26]int]bool{}
	var cur [26]int
	seen[cur] = true

	for j := 0; j < len(s2); j++ {
		cur[s2[j]-'a']++
		if seen[sub(cur, exp)] {
			return true
		}
		seen[cur] = true
	}

	return false
}

func sub(a [26]int, b [26]int) [26]int {
	var res [26]int
	for i := 0; i < 26; i++ {
		res[i] = a[i] - b[i]
	}
	return res
}
