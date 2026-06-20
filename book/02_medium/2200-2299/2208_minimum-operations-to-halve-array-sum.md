# 2208 — Minimum Operations To Halve Array Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func halveArray(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Heap

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Heap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2208: Minimum Operations to Halve Array Sum
// https://leetcode.com/problems/minimum-operations-to-halve-array-sum/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type maxHeap []float64

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x any)        { *h = append(*h, x.(float64)) }
func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func halveArray(nums []int) int {
	h := &maxHeap{}
	heap.Init(h)
	total := 0.0
	for _, v := range nums {
		total += float64(v)
  // Push ke priority queue
		heap.Push(h, float64(v))
	}

	target := total / 2.0
	ops := 0
	for total > target {
  // Pop dari priority queue
		largest := heap.Pop(h).(float64)
		half := largest / 2.0
		total -= half
  // Push ke priority queue
		heap.Push(h, half)
		ops++
	}
	return ops
}

func main() {
	// Test case 1
	fmt.Println(halveArray([]int{5, 19, 8, 1}))
	// Expected: 3

	// Test case 2
	fmt.Println(halveArray([]int{3, 8, 20}))
	// Expected: 3

	// Test case 3
	fmt.Println(halveArray([]int{1}))
	// Expected: 1
}
```
