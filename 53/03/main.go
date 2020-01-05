package main

import "fmt"

func main() {
	fmt.Println(freqAlphabets("10#11#12"))
}

func freqAlphabets(s string) string {
	var res []byte
	for i := 0; i < len(s); i++ {
		var d byte
		if i < len(s)-2 && s[i+2] == '#' {
			d = (s[i]-'0')*10 + (s[i+1] - '0')
			i += 2
		} else {
			d = (s[i] - '0')
		}
		res = append(res, d+'a'-1)
	}
	return string(res)
}
