package main

func main() {

}

type WordDictionary struct {
	root trieNode
}

type trieNode struct {
	children [26]*trieNode
	leaf     bool
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{}
}

/** Adds a word into the data structure. */
func (t *WordDictionary) AddWord(word string) {
	n := &t.root
	for i := 0; i < len(word); i++ {
		j := word[i] - 'a'
		if n.children[j] == nil {
			n.children[j] = &trieNode{}
		}
		n = n.children[j]
	}
	n.leaf = true
}

func (t *WordDictionary) find(word string, n *trieNode) bool {
	if len(word) == 0 {
		return n.leaf
	}

	c := word[0]
	if c == '.' {
		for j := range n.children {
			if n.children[j] != nil && t.find(word[1:], n.children[j]) {
				return true
			}
		}
		return false
	} else {
		j := c - 'a'
		if n.children[j] == nil {
			return false
		}
		return t.find(word[1:], n.children[j])
	}
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (t *WordDictionary) Search(word string) bool {
	return t.find(word, &t.root)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
