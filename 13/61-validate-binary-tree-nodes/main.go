package main

func main() {
}

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	notRoot := make([]bool, n)
	for _, n := range leftChild {
		if n != -1 {
			notRoot[n] = true
		}
	}
	for _, n := range rightChild {
		if n != -1 {
			notRoot[n] = true
		}
	}
	root := -1
	for i := range notRoot {
		if !notRoot[i] {
			if root != -1 {
				return false
			}
			root = i
		}
	}
	seen := make([]bool, n)
	var walk func(i int) bool
	walk = func(i int) bool {
		if i == -1 {
			return true
		}
		if seen[i] {
			return false
		}
		seen[i] = true
		return walk(leftChild[i]) && walk(rightChild[i])
	}
	if !walk(root) {
		return false
	}
	for _, s := range seen {
		if !s {
			return false
		}
	}
	return true
}
