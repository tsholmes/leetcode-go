package main

func main() {

}

type Vector2D struct {
	vals [][]int
}

func Constructor(v [][]int) Vector2D {
	for len(v) > 0 && len(v[0]) == 0 {
		v = v[1:]
	}
	return Vector2D{v}
}

func (v *Vector2D) Next() int {
	r := v.vals[0][0]
	v.vals[0] = v.vals[0][1:]
	for len(v.vals) > 0 && len(v.vals[0]) == 0 {
		v.vals = v.vals[1:]
	}
	return r
}

func (v *Vector2D) HasNext() bool {
	return len(v.vals) > 0
}

/**
 * Your Vector2D object will be instantiated and called as such:
 * obj := Constructor(v);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
