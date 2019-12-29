package main

import "math"

func main() {

}

type WordDistance struct {
	ws map[string][]int
}

func Constructor(words []string) WordDistance {
	ws := make(map[string][]int)
	for i, w := range words {
		ws[w] = append(ws[w], i)
	}
	return WordDistance{ws: ws}
}

func (wd *WordDistance) Shortest(word1 string, word2 string) int {
	w1, w2 := wd.ws[word1], wd.ws[word2]

	min := math.MaxInt32
	for _, i1 := range w1 {
		for _, i2 := range w2 {
			d := i1 - i2
			if d < 0 {
				d = -d
			}
			if d < min {
				min = d
			}
		}
	}
	return min
}

/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Shortest(word1,word2);
 */
