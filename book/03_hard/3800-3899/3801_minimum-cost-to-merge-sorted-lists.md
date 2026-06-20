# 3801 — Minimum Cost To Merge Sorted Lists

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minCost(lists [][]int) int
```

> **💡 Hint:** Always merge the two shortest lists (Huffman coding).

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS, Heap / Priority Queue, Stack, Merge Sort

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3801: Minimum Cost to Merge Sorted Lists
// https://leetcode.com/problems/minimum-cost-to-merge-sorted-lists/
// Difficulty: Hard
//
// Merge all sorted lists into one. Cost of each merge = sum of
// lengths of the two lists. Minimize total cost.
//
// Approach: Always merge the two shortest lists (Huffman coding).
// Use a min-heap (priority queue).

import (
	"container/heap"
	"fmt"
)

func main() {
	// Example 1
	fmt.Println(minCost([][]int{{1, 2}, {3, 4}, {5, 6}}))
	// Example 2
	fmt.Println(minCost([][]int{{1}, {2}, {3, 4}}))
	// Edge: single list
	fmt.Println(minCost([][]int{{1, 2, 3}}))
	// Edge: empty input
	fmt.Println(minCost([][]int{}))
}

func minCost(lists [][]int) int {
	if len(lists) <= 1 {
		return 0
	}

	h := &IntHeap{}
	heap.Init(h)
	for _, lst := range lists {
  // Masukkan elemen ke priority queue
		heap.Push(h, len(lst))
	}

	total := 0
	for h.Len() > 1 {
  // Ambil elemen terkecil/terbesar dari heap
		a := heap.Pop(h).(int)
  // Ambil elemen terkecil/terbesar dari heap
		b := heap.Pop(h).(int)
		cost := a + b
		total += cost
  // Masukkan elemen ke priority queue
		heap.Push(h, cost)
	}

	return total
}

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
```
