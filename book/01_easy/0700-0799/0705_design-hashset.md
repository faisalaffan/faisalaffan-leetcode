# 0705 — Design Hashset

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diminta untuk mendesain (merancang) sebuah struktur data kustom dengan operasi tertentu (insert, delete, search, update). Tugasmu adalah memilih representasi data yang tepat agar setiap operasi berjalan efisien — biasanya O(1) atau O(log n).

Ini adalah soal yang paling sering muncul di interview sistem desain. Kamu perlu memilih kombinasi struktur data yang tepat (HashMap + Heap + LinkedList) untuk mencapai kompleksitas yang diminta.

**Konsep kunci:** HashMap (O(1) lookup), Heap (priority), Doubly Linked List (O(1) remove), TreeMap (ordered keys).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor() MyHashSet
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1) average. Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #705: Design HashSet
// https://leetcode.com/problems/design-hashset/
// Difficulty: Easy

import "fmt"

const hashSetSize = 1000

type hashNode struct {
	key  int
	next *hashNode
}

// MyHashSet implements a hash set using separate chaining.
type MyHashSet struct {
	buckets []*hashNode
}

// Constructor creates a MyHashSet.
func Constructor() MyHashSet {
	return MyHashSet{buckets: make([]*hashNode, hashSetSize)}
}

// Add inserts a key into the set.
// Time: O(1) average. Space: O(n).
func (s *MyHashSet) Add(key int) {
	if s.Contains(key) {
		return
	}
	idx := key % hashSetSize
	s.buckets[idx] = &hashNode{key: key, next: s.buckets[idx]}
}

// Remove deletes a key from the set.
func (s *MyHashSet) Remove(key int) {
	idx := key % hashSetSize
	curr := s.buckets[idx]
	var prev *hashNode
	for curr != nil {
		if curr.key == key {
			if prev == nil {
				s.buckets[idx] = curr.next
			} else {
				prev.next = curr.next
			}
			return
		}
		prev = curr
		curr = curr.next
	}
}

// Contains checks if a key exists in the set.
func (s *MyHashSet) Contains(key int) bool {
	idx := key % hashSetSize
	curr := s.buckets[idx]
	for curr != nil {
		if curr.key == key {
			return true
		}
		curr = curr.next
	}
	return false
}

func main() {
	hs := Constructor()
	hs.Add(1)
	hs.Add(2)
	fmt.Println(hs.Contains(1)) // true
	fmt.Println(hs.Contains(3)) // false
	hs.Add(2)
	fmt.Println(hs.Contains(2)) // true
	hs.Remove(2)
	fmt.Println(hs.Contains(2)) // false
}
```
