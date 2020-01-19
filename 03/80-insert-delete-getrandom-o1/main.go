package main

import "math/rand"

func main() {

}

type RandomizedSet struct {
	idxs map[int]int
	keys []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{idxs: map[int]int{}}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.idxs[val]; ok {
		return false
	}
	r.idxs[val] = len(r.keys)
	r.keys = append(r.keys, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (r *RandomizedSet) Remove(val int) bool {
	if _, ok := r.idxs[val]; !ok {
		return false
	}

	i, j := r.idxs[val], len(r.keys)-1
	r.keys[i] = r.keys[j]
	r.idxs[r.keys[i]] = i
	r.keys = r.keys[:j]
	delete(r.idxs, val)
	return true
}

/** Get a random element from the set. */
func (r *RandomizedSet) GetRandom() int {
	if len(r.keys) == 0 {
		return 0
	}
	return r.keys[rand.Intn(len(r.keys))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
