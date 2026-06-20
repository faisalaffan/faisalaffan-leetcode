# 3275 — K Th Nearest Obstacle Queries

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func resultsArray(queries [][]int, k int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Heap / Priority Queue, Stack

**Kompleksitas Waktu:** O(n log k) Space: O(k)  
**Kompleksitas Ruang:** O(k)

> **Untuk fresh graduate:** Kuasai dulu teknik **Heap / Priority Queue** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3275: K-th Nearest Obstacle Queries
// https://leetcode.com/problems/k-th-nearest-obstacle-queries/
// Difficulty: Medium
// Time: O(n log k) Space: O(k)

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(resultsArray([][]int{{1, 2}, {3, 4}, {2, 3}, {-3, 0}}, 2)) // [-1,7,5,3]
	fmt.Println(resultsArray([][]int{{5, 5}, {4, 4}, {3, 3}}, 1))          // [10,8,6]
	fmt.Println(resultsArray([][]int{{1, 1}, {2, 2}, {3, 3}}, 3))          // [-1,-1,6]
}

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func resultsArray(queries [][]int, k int) []int {
  // Alokasi slice integer
	ans := make([]int, len(queries))
	h := &MaxHeap{}
	heap.Init(h)

	for i, q := range queries {
		dist := abs(q[0]) + abs(q[1])
  // Masukkan elemen ke priority queue
		heap.Push(h, dist)
		if h.Len() > k {
  // Ambil elemen terkecil/terbesar dari heap
			heap.Pop(h)
		}
		if h.Len() == k {
			ans[i] = (*h)[0]
		} else {
			ans[i] = -1
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
