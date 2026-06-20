# 1756 — Design Most Recently Used Queue

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diminta untuk mendesain (merancang) sebuah struktur data kustom dengan operasi tertentu (insert, delete, search, update). Tugasmu adalah memilih representasi data yang tepat agar setiap operasi berjalan efisien — biasanya O(1) atau O(log n).

Ini adalah soal yang paling sering muncul di interview sistem desain. Kamu perlu memilih kombinasi struktur data yang tepat (HashMap + Heap + LinkedList) untuk mencapai kompleksitas yang diminta.

**Konsep kunci:** HashMap (O(1) lookup), Heap (priority), Doubly Linked List (O(1) remove), TreeMap (ordered keys).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(n int) MRUQueue
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** O(n) per operation, Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1756: Design Most Recently Used Queue
// https://leetcode.com/problems/design-most-recently-used-queue/
// Difficulty: Medium [Paid]
// Time: O(n) per operation, Space: O(n)

import "fmt"

type MRUQueue struct {
	data []int
}

func Constructor(n int) MRUQueue {
  // Alokasi slice integer
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i + 1
	}
	return MRUQueue{data}
}

func (q *MRUQueue) Fetch(k int) int {
	// 1-indexed, fetch kth element and move to end
	val := q.data[k-1]
	// Remove
	q.data = append(q.data[:k-1], q.data[k:]...)
	// Move to end
	q.data = append(q.data, val)
	return val
}

func main() {
	q := Constructor(8)
	fmt.Println(q.Fetch(3)) // Expected: 3
	fmt.Println(q.Fetch(5)) // Expected: 6
	fmt.Println(q.Fetch(2)) // Expected: 2
	fmt.Println(q.Fetch(8)) // Expected: 3
}
```
