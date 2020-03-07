package main

import "strings"

func main() {

}

func licenseKeyFormatting(S string, K int) string {
	bs := []byte(S)
	count := 0
	for _, b := range bs {
		if b != '-' {
			bs[count] = b
			count++
		}
	}
	if count == 0 {
		return ""
	}
	S = strings.ToUpper(string(bs[:count]))
	first := len(S) % K
	if first == 0 {
		first = K
	}
	var res []byte
	res = append(res, S[:first]...)
	S = S[first:]
	for len(S) > 0 {
		res = append(res, '-')
		res = append(res, S[:K]...)
		S = S[K:]
	}
	return string(res)
}
