# 2599 — Make The Prefix Sum Non Negative

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func makePrefSumNonNegative(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Heap / Priority Queue, Stack, Prefix Sum

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Heap / Priority Queue** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2599: Make the Prefix Sum Non-negative
// https://leetcode.com/problems/make-the-prefix-sum-non-negative/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func makePrefSumNonNegative(nums []int) int {
	h := &MinHeap{}
	heap.Init(h)
	prefix := int64(0)
	moves := 0

	for _, v := range nums {
		prefix += int64(v)
		if v < 0 {
  // Masukkan elemen ke priority queue
			heap.Push(h, v)
		}
		for prefix < 0 {
			// Move the most negative seen to the end
  // Ambil elemen terkecil/terbesar dari heap
			smallest := heap.Pop(h).(int)
			prefix -= int64(smallest)
			moves++
		}
	}
	return moves
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", makePrefSumNonNegative([]int{2, 3, -5, 4}))
	// Expected: 0

	// Test case 2
	fmt.Println("Test 2:", makePrefSumNonNegative([]int{3, -5, -2, 6}))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", makePrefSumNonNegative([]int{1, -2, 3, -4, 5}))
	// Expected: 1
}
```
