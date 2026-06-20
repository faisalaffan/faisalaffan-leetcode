# 0706 — Design Hashmap

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diminta untuk mendesain (merancang) sebuah struktur data kustom dengan operasi tertentu (insert, delete, search, update). Tugasmu adalah memilih representasi data yang tepat agar setiap operasi berjalan efisien — biasanya O(1) atau O(log n).

Ini adalah soal yang paling sering muncul di interview sistem desain. Kamu perlu memilih kombinasi struktur data yang tepat (HashMap + Heap + LinkedList) untuk mencapai kompleksitas yang diminta.

**Konsep kunci:** HashMap (O(1) lookup), Heap (priority), Doubly Linked List (O(1) remove), TreeMap (ordered keys).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor() MyHashMap
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1) average. Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #706: Design HashMap
// https://leetcode.com/problems/design-hashmap/
// Difficulty: Easy

import "fmt"

const hashMapSize = 1000

type kvNode struct {
	key   int
	value int
	next  *kvNode
}

// MyHashMap implements a hash map using separate chaining.
type MyHashMap struct {
	buckets []*kvNode
}

// Constructor creates a MyHashMap.
func Constructor() MyHashMap {
	return MyHashMap{buckets: make([]*kvNode, hashMapSize)}
}

// Put inserts a key-value pair into the map.
// Time: O(1) average. Space: O(n).
func (m *MyHashMap) Put(key int, value int) {
	idx := key % hashMapSize
	curr := m.buckets[idx]
	for curr != nil {
		if curr.key == key {
			curr.value = value
			return
		}
		curr = curr.next
	}
	m.buckets[idx] = &kvNode{key: key, value: value, next: m.buckets[idx]}
}

// Get returns the value for a key, or -1 if not found.
func (m *MyHashMap) Get(key int) int {
	idx := key % hashMapSize
	curr := m.buckets[idx]
	for curr != nil {
		if curr.key == key {
			return curr.value
		}
		curr = curr.next
	}
	return -1
}

// Remove deletes a key from the map.
func (m *MyHashMap) Remove(key int) {
	idx := key % hashMapSize
	curr := m.buckets[idx]
	var prev *kvNode
	for curr != nil {
		if curr.key == key {
			if prev == nil {
				m.buckets[idx] = curr.next
			} else {
				prev.next = curr.next
			}
			return
		}
		prev = curr
		curr = curr.next
	}
}

func main() {
	hm := Constructor()
	hm.Put(1, 1)
	hm.Put(2, 2)
	fmt.Println(hm.Get(1)) // 1
	fmt.Println(hm.Get(3)) // -1
	hm.Put(2, 1)
	fmt.Println(hm.Get(2)) // 1
	hm.Remove(2)
	fmt.Println(hm.Get(2)) // -1
}
```
