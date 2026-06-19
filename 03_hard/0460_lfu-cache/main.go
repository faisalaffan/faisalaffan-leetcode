package main

import "fmt"

// LeetCode #460: LFU Cache
// https://leetcode.com/problems/lfu-cache/
// Difficulty: Hard
//
// Each frequency has a doubly linked LRU list. On access, move node to next
// frequency's list. On eviction, remove from the lowest frequency's LRU tail.
// get and put in O(1) amortized.

func main() {
	lfu := ConstructorLFU(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	fmt.Println("Get 1:", lfu.Get(1)) // 1
	lfu.Put(3, 3)                     // evicts key 2
	fmt.Println("Get 2:", lfu.Get(2)) // -1
	fmt.Println("Get 3:", lfu.Get(3)) // 3
	lfu.Put(4, 4)                     // evicts key 1
	fmt.Println("Get 1:", lfu.Get(1)) // -1
	fmt.Println("Get 3:", lfu.Get(3)) // 3
	fmt.Println("Get 4:", lfu.Get(4)) // 4

	// Single capacity
	lfu2 := ConstructorLFU(1)
	lfu2.Put(0, 0)
	fmt.Println("Get 0:", lfu2.Get(0)) // 0
	lfu2.Put(1, 1)
	fmt.Println("Get 0:", lfu2.Get(0)) // -1
	fmt.Println("Get 1:", lfu2.Get(1)) // 1
}

// LFUNode is a node in the LFU cache.
type LFUNode struct {
	key, value, freq int
	prev, next       *LFUNode
}

// FreqList is a doubly linked list for a specific frequency.
type FreqList struct {
	head, tail *LFUNode
}

func newFreqList() *FreqList {
	h := &LFUNode{}
	t := &LFUNode{}
	h.next = t
	t.prev = h
	return &FreqList{head: h, tail: t}
}

func (fl *FreqList) isEmpty() bool {
	return fl.head.next == fl.tail
}

// pushFront adds node to the front (most recently used).
func (fl *FreqList) pushFront(node *LFUNode) {
	node.prev = fl.head
	node.next = fl.head.next
	fl.head.next.prev = node
	fl.head.next = node
}

// remove removes a node from its current list.
func remove(node *LFUNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// popBack removes and returns the least recently used node.
func (fl *FreqList) popBack() *LFUNode {
	node := fl.tail.prev
	remove(node)
	return node
}

// LFUCache is the LFU cache.
type LFUCache struct {
	capacity int
	minFreq  int
	nodes    map[int]*LFUNode
	freqs    map[int]*FreqList
}

func ConstructorLFU(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		nodes:    make(map[int]*LFUNode),
		freqs:    make(map[int]*FreqList),
	}
}

func (c *LFUCache) Get(key int) int {
	node, ok := c.nodes[key]
	if !ok {
		return -1
	}
	c.incrementFreq(node)
	return node.value
}

func (c *LFUCache) Put(key int, value int) {
	if c.capacity == 0 {
		return
	}
	if node, ok := c.nodes[key]; ok {
		node.value = value
		c.incrementFreq(node)
		return
	}
	if len(c.nodes) >= c.capacity {
		c.evict()
	}
	node := &LFUNode{key: key, value: value, freq: 1}
	c.nodes[key] = node
	if c.freqs[1] == nil {
		c.freqs[1] = newFreqList()
	}
	c.freqs[1].pushFront(node)
	c.minFreq = 1
}

func (c *LFUCache) incrementFreq(node *LFUNode) {
	f := node.freq
	remove(node)
	if c.freqs[f].isEmpty() && f == c.minFreq {
		c.minFreq++
	}

	node.freq++
	if c.freqs[node.freq] == nil {
		c.freqs[node.freq] = newFreqList()
	}
	c.freqs[node.freq].pushFront(node)
}

func (c *LFUCache) evict() {
	list := c.freqs[c.minFreq]
	if list == nil || list.isEmpty() {
		return
	}
	node := list.popBack()
	delete(c.nodes, node.key)
}
