package main

func main() {
	l := Constructor(1)
	// l.Put(1, 1)
	// l.Put(2, 2)
	// fmt.Println(l.Get(1))
	// l.Put(3, 3)
	// fmt.Println(l.Get(2))
	// fmt.Println(l.Get(3))
	// l.Put(4, 4)
	// fmt.Println(l.Get(1))
	// fmt.Println(l.Get(3))
	// fmt.Println(l.Get(4))

	l.Put(2, 1)
	l.Get(2)
	l.Put(3, 2)
	l.Get(2)
	l.Get(3)
}

type LFUCache struct {
	vals   map[int]int
	counts map[int][2]int
	iter   int
	cap    int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		vals:   map[int]int{},
		counts: map[int][2]int{},
		cap:    capacity,
	}
}

func (l *LFUCache) Get(key int) int {
	if v, ok := l.vals[key]; ok {
		l.bump(key)
		return v
	}
	return -1
}

func (l *LFUCache) Put(key int, value int) {
	if l.cap <= 0 {
		return
	}

	if _, ok := l.vals[key]; ok {
		l.vals[key] = value
		l.bump(key)
		return
	}

	if len(l.counts) == l.cap {
		mink := -1
		var min [2]int
		for k, v := range l.counts {
			if mink == -1 || (v[1] < min[1] || (v[1] == min[1] && v[0] < min[0])) {
				mink = k
				min = v
			}
		}
		delete(l.vals, mink)
		delete(l.counts, mink)
	}

	l.vals[key] = value
	l.bump(key)
}

func (l *LFUCache) bump(key int) {
	it := l.iter
	l.iter++
	c := l.counts[key]
	c = [2]int{it, c[1] + 1}
	l.counts[key] = c
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
