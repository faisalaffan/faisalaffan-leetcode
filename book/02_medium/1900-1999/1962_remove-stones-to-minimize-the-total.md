# 1962 — Remove Stones To Minimize The Total

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MinStoneSum(piles []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Heap

**Waktu:** O(n + k log n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Heap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1962: Remove Stones to Minimize the Total
// https://leetcode.com/problems/remove-stones-to-minimize-the-total/
// Difficulty: Medium

import (
	"container/heap"
	"fmt"
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
	fmt.Println(MinStoneSum([]int{5, 4, 9}, 2))
	fmt.Println(MinStoneSum([]int{4, 3, 6, 7}, 3))
}

// Time: O(n + k log n), Space: O(n)
func MinStoneSum(piles []int, k int) int {
	h := &MaxHeap{}
	heap.Init(h)
	sum := 0
	for _, p := range piles {
		sum += p
  // Push ke priority queue
		heap.Push(h, p)
	}

	for i := 0; i < k; i++ {
  // Pop dari priority queue
		cur := heap.Pop(h).(int)
		removed := cur / 2
		sum -= removed
  // Push ke priority queue
		heap.Push(h, cur-removed)
	}
	return sum
}
```
