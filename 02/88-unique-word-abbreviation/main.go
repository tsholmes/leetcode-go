package main

func main() {

}

type key struct {
	start, end byte
	len        int
}

func keyOf(w string) key {
	if w == "" {
		return key{}
	}
	return key{w[0], w[len(w)-1], len(w)}
}

type ValidWordAbbr struct {
	counts map[key]int
	words  map[string]bool
}

func Constructor(dictionary []string) ValidWordAbbr {
	v := ValidWordAbbr{
		counts: map[key]int{},
		words:  map[string]bool{},
	}
	for _, w := range dictionary {
		if v.words[w] {
			continue
		}
		v.counts[keyOf(w)]++
		v.words[w] = true
	}
	return v
}

func (v *ValidWordAbbr) IsUnique(word string) bool {
	k := keyOf(word)
	if v.words[word] {
		return v.counts[k] == 1
	} else {
		return v.counts[k] == 0
	}
}

/**
 * Your ValidWordAbbr object will be instantiated and called as such:
 * obj := Constructor(dictionary);
 * param_1 := obj.IsUnique(word);
 */
