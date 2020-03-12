package main

import "strconv"

func main() {

}

func wordsAbbreviation(dict []string) []string {
	type trieKey struct {
		c byte
		l byte
		r int
	}
	type trieNode struct {
		count int
		cs    map[trieKey]*trieNode
	}
	var add func(n *trieNode, w string)
	add = func(n *trieNode, w string) {
		n.count++
		if len(w) == 0 {
			return
		}
		k := trieKey{
			c: w[0],
			l: w[len(w)-1],
			r: len(w),
		}
		if n.cs[k] == nil {
			n.cs[k] = &trieNode{cs: map[trieKey]*trieNode{}}
		}
		add(n.cs[k], w[1:])
	}
	root := &trieNode{cs: map[trieKey]*trieNode{}}
	for _, w := range dict {
		add(root, w)
	}

	res := make([]string, len(dict))
	var work []byte
	for i, w := range dict {
		l := 1
		ec := w[len(w)-1]
		n := root.cs[trieKey{
			c: w[0],
			l: ec,
			r: len(w),
		}]
		for n.count > 1 && l < len(w) {
			n = n.cs[trieKey{
				c: w[l],
				l: ec,
				r: len(w) - l,
			}]
			l++
		}
		rem := len(w) - l - 1
		if rem < 2 {
			res[i] = w
		} else {
			work = append(work[:0], w[:l]...)
			work = strconv.AppendInt(work, int64(rem), 10)
			work = append(work, w[len(w)-1])
			res[i] = string(work)
		}
	}

	return res
}
