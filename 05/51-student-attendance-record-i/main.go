package main

func main() {

}

func checkRecord(s string) bool {
	acount := 0
	lcount := 0

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'P':
			lcount = 0
		case 'A':
			acount++
			lcount = 0
		case 'L':
			lcount++
			if lcount > 2 {
				return false
			}
		}
	}

	return acount <= 1
}
