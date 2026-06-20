# 0146 — Lru Cache

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func Constructor(capacity int) LRUCache`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(1) per operation  |  **Ruang:** O(capacity)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #146: LRU Cache
// https://leetcode.com/problems/lru-cache/
// Difficulty: Medium

import "fmt"

type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node // LRU
	tail     *Node // MRU
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToFront(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.moveToFront(node)
		return
	}

	if len(this.cache) == this.capacity {
		this.removeLRU()
	}

	node := &Node{key: key, value: value}
	this.cache[key] = node
	this.addToFront(node)
}

func (this *LRUCache) moveToFront(node *Node) {
	this.removeNode(node)
	this.addToFront(node)
}

func (this *LRUCache) addToFront(node *Node) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) removeLRU() {
	lru := this.tail.prev
	this.removeNode(lru)
	delete(this.cache, lru.key)
}

func main() {
	// Test case 1
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1)) // 1
	lru.Put(3, 3)           // evicts key 2
	fmt.Println(lru.Get(2)) // -1
	lru.Put(4, 4)           // evicts key 1
	fmt.Println(lru.Get(1)) // -1
	fmt.Println(lru.Get(3)) // 3
	fmt.Println(lru.Get(4)) // 4
}

// Time: O(1) per operation | Space: O(capacity)
```
