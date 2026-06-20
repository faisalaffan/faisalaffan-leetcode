# 3362 — Zero Array Transformation Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxRemoval(nums []int, queries [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Heap / Priority Queue, Stack

**Kompleksitas Waktu:** O(n log n) Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3362: Zero Array Transformation III
// https://leetcode.com/problems/zero-array-transformation-iii/
// Difficulty: Medium
// Time: O(n log n) Space: O(n)

import (
	"container/heap"
	"fmt"
	"sort"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	fmt.Println(maxRemoval([]int{2, 0, 2}, [][]int{{0, 2}, {0, 2}, {1, 1}})) // 1
}

func maxRemoval(nums []int, queries [][]int) int {
	n := len(nums)
	m := len(queries)

	// Sort queries by left endpoint
  // Custom sort dengan comparator
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][0] < queries[j][0]
	})

	// For each position i, we need nums[i] decrements at i
	// Greedily use queries with farthest right endpoint
	h := &MaxHeap{}
	heap.Init(h)

  // Alokasi slice integer
	diff := make([]int, n+1)
	qi := 0
	cur := 0

	for i := 0; i < n; i++ {
		// Add all queries that start at i
		for qi < m && queries[qi][0] == i {
  // Masukkan elemen ke priority queue
			heap.Push(h, queries[qi][1])
			qi++
		}

		cur += diff[i]
		need := nums[i] - cur

		for need > 0 && h.Len() > 0 {
  // Ambil elemen terkecil/terbesar dari heap
			r := heap.Pop(h).(int)
			cur++
			diff[r+1]++
			need--
		}

		if need > 0 {
			return -1
		}
	}

	// Unused queries can be removed
	return h.Len()
}
```
