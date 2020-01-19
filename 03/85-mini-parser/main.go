package main

func main() {

}

func deserialize(s string) *NestedInteger {
	i := 0
	var build func() *NestedInteger
	build = func() *NestedInteger {
		var n NestedInteger
		if s[i] == '[' {
			// list
			i++
			for s[i] != ']' {
				if s[i] == ',' {
					i++
				}
				n.Add(*build())
			}
			i++
		} else {
			// num
			neg := false
			v := 0
			if s[i] == '-' {
				neg = true
				i++
			}
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				v = v*10 + int(s[i]-'0')
				i++
			}
			if neg {
				v = -v
			}
			n.SetInteger(v)
		}
		return &n
	}
	return build()
}

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	isInt    bool
	val      int
	children []*NestedInteger
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool { return n.isInt }

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int { return n.val }

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) { n.isInt, n.val, n.children = true, value, nil }

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {
	n.isInt, n.val, n.children = false, 0, append(n.children, &elem)
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger { return n.children }
