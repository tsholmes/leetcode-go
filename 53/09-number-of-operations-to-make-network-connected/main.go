package main

func main() {

}
func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}
	ds := newDisjointSubset(n)
	for _, con := range connections {
		ds.join(con[0], con[1])
	}
	return ds.countRoots() - 1
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
