package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(minAbbreviation("apple", []string{"plain", "amber", "blade"}))

	target := "aaaaaaaaaaaaaaaaaaaa"
	dict := make([]string, 1000)
	for i := range dict {
		w := []byte(target)
		for j := range w {
			if i&(1<<uint(j)) == 0 {
				w[j] = 'b'
			}
		}
		dict[i] = string(w)
	}
	fmt.Println(minAbbreviation(target, dict))
}

func minAbbreviation(target string, dictionary []string) string {
	type trieNode struct {
		letter *trieNode
		counts [20]*trieNode
		leaf   bool
	}

	var insert func(n *trieNode, i int, w string, lastn bool)
	insert = func(n *trieNode, i int, w string, lastn bool) {
		if len(w) == 0 {
			n.leaf = true
			return
		}
		if w[0] == target[i] {
			if n.letter == nil {
				n.letter = &trieNode{}
			}
			insert(n.letter, i+1, w[1:], false)
		}
		if !lastn {
			for c := 0; c+2 <= len(w); c++ {
				if n.counts[c] == nil {
					n.counts[c] = &trieNode{}
				}
				insert(n.counts[c], i+c+2, w[c+2:], true)
			}
		}
	}

	root := &trieNode{}
	for _, w := range dictionary {
		if len(w) == len(target) {
			insert(root, 0, w, false)
		}
	}

	minw := target
	minp := len(target)
	var work []byte
	var search func(w string, plen int, n *trieNode, lastn bool)
	search = func(w string, plen int, n *trieNode, lastn bool) {
		if plen >= minp {
			return
		}
		if len(w) > 0 && plen+1 >= minp {
			return
		}
		if len(w) == 0 {
			if n == nil {
				minw = string(work)
				minp = plen
			}
			return
		}

		wlen := len(work)

		if n == nil {
			work = strconv.AppendInt(work, int64(len(w)), 10)
			search("", plen+1, nil, true)
			work = work[:wlen]
			return
		}

		if !lastn {
			for c := len(w) - 2; c >= 0; c-- {
				work = strconv.AppendInt(work, int64(c+2), 10)
				search(w[c+2:], plen+1, n.counts[c], true)
				work = work[:wlen]
			}
		}

		work = append(work, w[0])
		search(w[1:], plen+1, n.letter, false)
		work = work[:wlen]
	}

	search(target, 0, root, false)

	return minw
}
