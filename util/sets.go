package util

// This is meant to be copied into a solution

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

// optional util methods

func (d disjointSubset) isSingle() bool {
	for i := range d {
		if d.parent(i) != d.parent(0) {
			return false
		}
	}
	return true
}

func (d disjointSubset) getSiblings(n int) []int {
	var res []int
	p := d.parent(n)
	for i := range d {
		if d.parent(i) == p {
			res = append(res, i)
		}
	}
	return res
}
