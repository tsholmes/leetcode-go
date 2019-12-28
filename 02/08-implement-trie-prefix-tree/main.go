package main

func main() {

}

type Trie struct {
	root trieNode
}

type trieNode struct {
	children [26]*trieNode
	leaf     bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
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

func (t *Trie) find(word string) *trieNode {
	n := &t.root
	for i := 0; i < len(word); i++ {
		j := word[i] - 'a'
		if n.children[j] == nil {
			return nil
		}
		n = n.children[j]
	}
	return n
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	n := t.find(word)
	return n != nil && n.leaf
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	return t.find(prefix) != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
