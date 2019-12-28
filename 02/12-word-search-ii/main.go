package main

import "fmt"

func main() {
	fmt.Println(findWords([][]byte{
		[]byte(`oaan`),
		[]byte(`etae`),
		[]byte(`ihkr`),
		[]byte(`iflv`),
	}, []string{"oath", "pea", "eat", "rain"}))
}

func findWords(board [][]byte, words []string) []string {
	if len(board) == 0 || len(board[0]) == 0 {
		return nil
	}

	type trieNode struct {
		children [26]*trieNode
		leaf     bool
		word     string
	}
	root := &trieNode{}
	for _, w := range words {
		n := root
		for i := 0; i < len(w); i++ {
			j := w[i] - 'a'
			if n.children[j] == nil {
				n.children[j] = &trieNode{}
			}
			n = n.children[j]
		}
		n.leaf = true
		n.word = w
	}

	hei, wid := len(board), len(board[0])
	dirs := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	found := map[string]bool{}
	used := map[[2]int]bool{}
	var search func([2]int, *trieNode)
	search = func(pos [2]int, n *trieNode) {
		if n == nil || used[pos] {
			return
		}

		if n.leaf {
			found[n.word] = true
		}

		used[pos] = true

		for _, dir := range dirs {
			p2 := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
			if p2[0] < 0 || p2[0] >= hei || p2[1] < 0 || p2[1] >= wid {
				continue
			}
			j := board[p2[0]][p2[1]] - 'a'
			search(p2, n.children[j])
		}

		used[pos] = false
	}

	for i := 0; i < hei; i++ {
		for j := 0; j < wid; j++ {
			search([2]int{i, j}, root.children[board[i][j]-'a'])
		}
	}

	var res []string
	for k := range found {
		res = append(res, k)
	}

	return res
}
