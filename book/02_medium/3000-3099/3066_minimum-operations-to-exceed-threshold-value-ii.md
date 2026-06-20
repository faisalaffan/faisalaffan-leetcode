# 3066 — Minimum Operations To Exceed Threshold Value Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minOperations3066(nums []int, k int) (ans int)`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Heap

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Heap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3066: Minimum Operations to Exceed Threshold Value II
// https://leetcode.com/problems/minimum-operations-to-exceed-threshold-value-ii/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	fmt.Println(minOperations3066([]int{2, 11, 10, 1, 3}, 10))
	fmt.Println(minOperations3066([]int{1, 1, 2, 4, 9}, 20))
}

func minOperations3066(nums []int, k int) (ans int) {
	h := &minHeap{}
	heap.Init(h)
	for _, x := range nums {
		if x < k {
  // Push ke priority queue
			heap.Push(h, x)
		}
	}
	for h.Len() >= 2 {
  // Pop dari priority queue
		x := heap.Pop(h).(int)
  // Pop dari priority queue
		y := heap.Pop(h).(int)
		val := x*2 + y
		if val < k {
  // Push ke priority queue
			heap.Push(h, val)
		}
		ans++
	}
	if h.Len() > 0 {
		ans++
	}
	return
}
```
