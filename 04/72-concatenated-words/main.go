package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(findAllConcatenatedWordsInADict([]string{
		"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat",
	}))
}

func findAllConcatenatedWordsInADict(words []string) []string {
	root := &trieNode{}
	for _, w := range words {
		if w == "" {
			continue
		}
		add(root, w, w)
	}

	type state struct {
		fn    *trieNode
		rn    *trieNode
		multi bool
	}
	queue := []state{
		{root, root, false},
	}
	seen := map[state]bool{}
	valid := map[string]bool{}
	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]
		if seen[s] {
			continue
		}
		seen[s] = true
		if s.multi && len(s.fn.fw) != 0 && len(s.rn.fw) != 0 {
			valid[s.fn.fw] = true
		}
		mask := s.fn.mask & s.rn.mask
		for mask != 0 {
			b := mask & -mask
			i := 63 - bits.LeadingZeros(uint(b))
			queue = append(queue, state{s.fn.cs[i], s.rn.cs[i], s.multi})
			if len(s.rn.cs[i].fw) != 0 {
				queue = append(queue, state{s.fn.cs[i], root, true})
			}
			mask = mask &^ b
		}
	}

	var res []string
	for k := range valid {
		res = append(res, k)
	}
	return res
}

type trieNode struct {
	cs   [26]*trieNode
	mask int
	fw   string
}

func add(n *trieNode, w string, fw string) {
	if len(w) == 0 {
		n.fw = fw
		return
	}
	c := w[0] - 'a'
	if n.cs[c] == nil {
		n.cs[c] = &trieNode{}
		n.mask |= 1 << uint(c)
	}
	add(n.cs[c], w[1:], fw)
}
