package main

import "fmt"

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // returns 1
	cache.Put(3, 3)           // evicts key 2
	fmt.Println(cache.Get(2)) // returns -1 (not found)
	cache.Put(4, 4)           // evicts key 1
	fmt.Println(cache.Get(1)) // returns -1 (not found)
	fmt.Println(cache.Get(3)) // returns 3
	fmt.Println(cache.Get(4)) // returns 4
}

type LRUCache struct {
	capacity int
	head     *lruItem
	tail     *lruItem
	items    map[int]*lruItem
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		items:    make(map[int]*lruItem, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	item := this.items[key]
	if item == nil {
		return -1
	}
	this.moveToHead(item)
	return item.val
}

func (this *LRUCache) Put(key int, value int) {
	item := this.items[key]
	if item != nil {
		item.val = value
		this.moveToHead(item)
		return
	}
	item = &lruItem{
		key:  key,
		val:  value,
		next: this.head,
	}
	if this.head != nil {
		this.head.prev = item
	} else {
		this.tail = item
	}
	this.head = item
	this.items[key] = item

	for len(this.items) > this.capacity {
		tail := this.tail
		this.tail = tail.prev
		delete(this.items, tail.key)
	}
}

func (this *LRUCache) moveToHead(item *lruItem) {
	if item == this.head {
		return
	}
	if item == this.tail && this.capacity > 1 {
		this.tail = item.prev
	}
	if item.prev != nil {
		item.prev.next = item.next
	}
	if item.next != nil {
		item.next.prev = item.prev
	}
	item.prev = nil
	item.next = this.head
	if this.head != nil {
		this.head.prev = item
	}
	this.head = item
}

type lruItem struct {
	key  int
	val  int
	prev *lruItem
	next *lruItem
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
