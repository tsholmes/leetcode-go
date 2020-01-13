package main

import (
	"fmt"
	"sort"
)

func main() {
	h := Constructor()
	h.Hit(1)
	h.Hit(2)
	h.Hit(3)
	fmt.Println(h.GetHits(4))
	h.Hit(300)
	fmt.Println(h.GetHits(300))
	fmt.Println(h.GetHits(301))
}

type HitCounter struct {
	hits []int
}

/** Initialize your data structure here. */
func Constructor() HitCounter {
	return HitCounter{}
}

/** Record a hit.
  @param timestamp - The current timestamp (in seconds granularity). */
func (h *HitCounter) Hit(timestamp int) {
	h.hits = append(h.hits, timestamp)
}

func (h *HitCounter) hitsUpTo(timestamp int) int {
	idx := sort.Search(len(h.hits), func(i int) bool { return h.hits[i] > timestamp })
	return idx
}

/** Return the number of hits in the past 5 minutes.
  @param timestamp - The current timestamp (in seconds granularity). */
func (h *HitCounter) GetHits(timestamp int) int {
	return h.hitsUpTo(timestamp) - h.hitsUpTo(timestamp-300)
}

/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */
