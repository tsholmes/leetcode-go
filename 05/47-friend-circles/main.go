package main

func main() {

}

func findCircleNum(M [][]int) int {
	N := len(M)
	dsj := newDisjointSubset(N)
	for i := 1; i < N; i++ {
		for j := 0; j < i; j++ {
			if M[i][j] == 1 {
				dsj.join(i, j)
			}
		}
	}
	return dsj.countRoots()
}

type disjointSubset []int

func newDisjointSubset(n int) disjointSubset {
	res := make(disjointSubset, n)
	for i := range res {
		res[i] = i
	}
	return res
}

func (d disjointSubset) parent(n int) int {
	if d[n] == n {
		return n
	}
	p := d.parent(d[n])
	d[n] = p
	return p
}

func (d disjointSubset) join(a int, b int) {
	pa, pb := d.parent(a), d.parent(b)
	d[pa] = pb
}

func (d disjointSubset) countRoots() int {
	count := 0
	for i := range d {
		if d.parent(i) == i {
			count++
		}
	}
	return count
}
