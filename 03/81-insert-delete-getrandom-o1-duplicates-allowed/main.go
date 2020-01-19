package main

import (
	"fmt"
	"math/rand"
)

func main() {
	r := Constructor()
	r.Insert(0)
	fmt.Println(r)
	r.Insert(1)
	fmt.Println(r)
	r.Remove(0)
	fmt.Println(r)
	r.Insert(2)
	fmt.Println(r)
	r.Remove(1)
	fmt.Println(r)
	fmt.Println(r.GetRandom())
}

type RandomizedCollection struct {
	idxs map[int][]int
	keys [][2]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{idxs: map[int][]int{}}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (r *RandomizedCollection) Insert(val int) bool {
	ex, ok := r.idxs[val]

	i := len(r.keys)
	c := len(ex)

	r.idxs[val] = append(ex, i)
	r.keys = append(r.keys, [2]int{val, c})

	return !ok
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (r *RandomizedCollection) Remove(val int) bool {
	idxs := r.idxs[val]
	if len(idxs) == 0 {
		return false
	}

	i, j := idxs[len(idxs)-1], len(r.keys)-1
	r.keys[i] = r.keys[j]
	r.idxs[r.keys[i][0]][r.keys[i][1]] = i

	r.keys = r.keys[:j]

	if len(idxs) > 1 {
		r.idxs[val] = idxs[:len(idxs)-1]
	} else {
		delete(r.idxs, val)
	}

	return true
}

/** Get a random element from the collection. */
func (r *RandomizedCollection) GetRandom() int {
	return r.keys[rand.Intn(len(r.keys))][0]
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
