package main

import "fmt"

func main() {
	fmt.Println(wordSquares([]string{"area", "lead", "wall", "lady", "ball"}))
}

func wordSquares(words []string) [][]string {
	if len(words) == 0 {
		return nil
	}
	N := len(words[0])

	type trieNode struct {
		cs [26]*trieNode
		ws []string
	}

	root := &trieNode{}
	for _, w := range words {
		n := root
		for i := 0; i < N; i++ {
			ix := int(w[i] - 'a')
			if n.cs[ix] == nil {
				n.cs[ix] = &trieNode{}
			}
			n = n.cs[ix]
			n.ws = append(n.ws, w)
		}
	}
	root.ws = words

	var res [][]string
	var search func(int, []*trieNode)
	var work []string
	search = func(i int, ns []*trieNode) {
		for _, n := range ns {
			if n == nil {
				return
			}
		}
		if i == N {
			res = append(res, append([]string{}, work...))
			return
		}

		var pns [5]*trieNode
		for _, w := range ns[i].ws {
			copy(pns[:], ns)

			work = append(work, w)

			for j := i; j < N; j++ {
				ns[j] = ns[j].cs[w[j]-'a']
			}
			search(i+1, ns)

			copy(ns, pns[:])
			work = work[:i]
		}

	}

	ns := make([]*trieNode, N)
	for i := range ns {
		ns[i] = root
	}

	search(0, ns)

	return res
}
