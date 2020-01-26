package main

import "fmt"

func main() {
	a := Constructor()
	a.Inc("a")
	fmt.Println(a.GetMinKey(), a.GetMaxKey(), a)
	a.Inc("a")
	fmt.Println(a.GetMinKey(), a.GetMaxKey(), a)
	a.Inc("b")
	fmt.Println(a.GetMinKey(), a.GetMaxKey(), a)
	a.Dec("b")
	fmt.Println(a.GetMinKey(), a.GetMaxKey(), a)
	a.Dec("a")
	a.Dec("a")
	fmt.Println(a.GetMinKey(), a.GetMaxKey(), a)
}

type AllOne struct {
	vals   map[string]int
	levels map[int]map[string]struct{}

	max int
	min int
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	return AllOne{
		vals:   map[string]int{},
		levels: map[int]map[string]struct{}{},
	}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (a *AllOne) Inc(key string) {
	cl := a.vals[key]
	if v, ok := a.levels[cl]; ok {
		delete(v, key)
	}

	if cl == a.min && len(a.levels[a.min]) == 0 {
		a.min++
	} else if cl == 0 {
		a.min = 1
	}
	if cl == a.max {
		a.max++
	}

	a.vals[key] = cl + 1
	if a.levels[cl+1] == nil {
		a.levels[cl+1] = map[string]struct{}{}
	}
	a.levels[cl+1][key] = struct{}{}
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (a *AllOne) Dec(key string) {
	cl := a.vals[key]
	if cl == 0 {
		return
	}
	if v, ok := a.levels[cl]; ok {
		delete(v, key)
	}

	if cl == a.min {
		a.min--
		if a.min == 0 {
			a.min = a.max
			for k, v := range a.levels {
				if len(v) > 0 && k < a.min {
					a.min = k
				}
			}
		}
	}
	if cl == a.max && len(a.levels[a.max]) == 0 {
		a.max--
	}

	if cl == 1 {
		delete(a.vals, key)
		return
	}
	a.vals[key] = cl - 1
	if a.levels[cl-1] == nil {
		a.levels[cl-1] = map[string]struct{}{}
	}
	a.levels[cl-1][key] = struct{}{}
}

/** Returns one of the keys with maximal value. */
func (a *AllOne) GetMaxKey() string {
	if a.max == 0 {
		return ""
	}
	for k := range a.levels[a.max] {
		return k
	}
	return ""
}

/** Returns one of the keys with Minimal value. */
func (a *AllOne) GetMinKey() string {
	if a.min == 0 {
		return ""
	}
	for k := range a.levels[a.min] {
		return k
	}
	return ""
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
