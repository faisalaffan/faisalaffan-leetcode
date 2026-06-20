# 1962 — Remove Stones To Minimize The Total

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinStoneSum(piles []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Heap / Priority Queue, Stack

**Kompleksitas Waktu:** O(n + k log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Heap / Priority Queue** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Masukkan elemen ke priority queue
		heap.Push(h, p)
	}

	for i := 0; i < k; i++ {
  // Ambil elemen terkecil/terbesar dari heap
		cur := heap.Pop(h).(int)
		removed := cur / 2
		sum -= removed
  // Masukkan elemen ke priority queue
		heap.Push(h, cur-removed)
	}
	return sum
}
```
