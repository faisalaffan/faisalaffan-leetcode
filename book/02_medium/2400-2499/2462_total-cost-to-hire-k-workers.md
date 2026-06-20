# 2462 — Total Cost To Hire K Workers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func totalCost(costs []int, k int, candidates int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Heap / Priority Queue, Stack

**Kompleksitas Waktu:** O((candidates + k) log candidates)  
**Kompleksitas Ruang:** O(candidates)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2462: Total Cost to Hire K Workers
// https://leetcode.com/problems/total-cost-to-hire-k-workers/
// Difficulty: Medium
// Time: O((candidates + k) log candidates) | Space: O(candidates)
// Two min-heaps: left and right. Each round pick cheaper candidate.

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

func main() {
	fmt.Println(totalCost([]int{17, 12, 10, 2, 7, 2, 11, 20, 8}, 3, 4)) // 11
	fmt.Println(totalCost([]int{1, 2, 4, 1}, 3, 3))                      // 4
}

func totalCost(costs []int, k int, candidates int) int64 {
	n := len(costs)
	leftH := &MinHeap{}
	rightH := &MinHeap{}
	heap.Init(leftH)
	heap.Init(rightH)

	left, right := 0, n-1
	for i := 0; i < candidates && left <= right; i++ {
  // Masukkan elemen ke priority queue
		heap.Push(leftH, costs[left])
		left++
	}
	for i := 0; i < candidates && left <= right; i++ {
  // Masukkan elemen ke priority queue
		heap.Push(rightH, costs[right])
		right--
	}

	var total int64
	for i := 0; i < k; i++ {
		if rightH.Len() == 0 || (leftH.Len() > 0 && (*leftH)[0] <= (*rightH)[0]) {
  // Ambil elemen terkecil/terbesar dari heap
			total += int64(heap.Pop(leftH).(int))
			if left <= right {
  // Masukkan elemen ke priority queue
				heap.Push(leftH, costs[left])
				left++
			}
		} else {
  // Ambil elemen terkecil/terbesar dari heap
			total += int64(heap.Pop(rightH).(int))
			if left <= right {
  // Masukkan elemen ke priority queue
				heap.Push(rightH, costs[right])
				right--
			}
		}
	}
	return total
}
```
