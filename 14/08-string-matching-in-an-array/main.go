package main

import "strings"

func main() {

}

func stringMatching(words []string) []string {
	var res []string
	for i, w1 := range words {
		for j, w2 := range words {
			if i == j {
				continue
			}
			if strings.Contains(w2, w1) {
				res = append(res, w1)
				break
			}
		}
	}
	return res
}
