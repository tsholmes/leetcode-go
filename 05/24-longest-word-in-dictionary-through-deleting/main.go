package main

import "sort"

func main() {

}

func findLongestWord(s string, d []string) string {
	sort.Slice(d, func(i, j int) bool {
		si, sj := d[i], d[j]
		if len(si) > len(sj) {
			return true
		} else if len(si) < len(sj) {
			return false
		}
		return si < sj
	})

	for _, w := range d {
		if isSubSeq(w, s) {
			return w
		}
	}
	return ""
}

func isSubSeq(a string, b string) bool {
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
			j++
		} else {
			j++
		}
	}
	return i == len(a)
}
