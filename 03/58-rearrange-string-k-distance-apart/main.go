package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(rearrangeString("aabbcc", 3))
	fmt.Println(rearrangeString("aaabc", 3))
	fmt.Println(rearrangeString("aaadbbcc", 2))
	fmt.Println(rearrangeString("aaabc", 2))
}

func rearrangeString(s string, k int) string {
	if k < 2 {
		return s
	}
	counts := map[byte]int{}
	for i := 0; i < len(s); i++ {
		counts[s[i]]++
	}

	var h charHeap
	for c, v := range counts {
		h = append(h, ccount{c, v})
	}

	heap.Init(&h)

	var res []byte
	var cooldown []ccount
	for i := 0; i < len(s); i++ {
		if h.Len() == 0 {
			return ""
		}
		cc := heap.Pop(&h).(ccount)
		res = append(res, cc.c)
		cc.count--
		cooldown = append(cooldown, cc)
		if len(cooldown) == k {
			cc, cooldown = cooldown[0], cooldown[1:]
			if cc.count > 0 {
				heap.Push(&h, cc)
			}
		}
	}

	return string(res)
}

type ccount struct {
	c     byte
	count int
}

type charHeap []ccount

func (h *charHeap) Len() int           { return len(*h) }
func (h *charHeap) Less(i, j int) bool { return (*h)[i].count > (*h)[j].count }
func (h *charHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *charHeap) Push(x interface{}) { *h = append(*h, x.(ccount)) }
func (h *charHeap) Pop() interface{}   { v := (*h)[h.Len()-1]; *h = (*h)[:h.Len()-1]; return v }
