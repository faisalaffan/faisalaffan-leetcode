# 2163 — Minimum Difference In Sums After Removal Of Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minimumDifference(nums []int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Heap, Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Heap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2163: Minimum Difference in Sums After Removal of Elements
// https://leetcode.com/problems/minimum-difference-in-sums-after-removal-of-elements/
// Difficulty: Hard
//
// Prefix min-heap (keep n smallest) and suffix max-heap (keep n largest).
// Answer = min over split points of prefixMin[i] - suffixMax[i+1].

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(minimumDifference([]int{3, 1, 2}))                       // -1
	fmt.Println(minimumDifference([]int{7, 9, 5, 8, 1, 3}))             // 1
	fmt.Println(minimumDifference([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))    // -18
}

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

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minimumDifference(nums []int) int64 {
	m := len(nums)
	n := m / 3

  // Alokasi slice
	pref := make([]int, m)
	h := &MaxHeap{}
	heap.Init(h)
	sum := 0
	for i := 0; i < m; i++ {
  // Push ke priority queue
		heap.Push(h, nums[i])
		sum += nums[i]
		if h.Len() > n {
  // Pop dari priority queue
			sum -= heap.Pop(h).(int)
		}
		if i >= n-1 {
			pref[i] = sum
		}
	}

  // Alokasi slice
	suf := make([]int, m)
	h2 := &MinHeap{}
	heap.Init(h2)
	sum = 0
	for i := m - 1; i >= 0; i-- {
  // Push ke priority queue
		heap.Push(h2, nums[i])
		sum += nums[i]
		if h2.Len() > n {
  // Pop dari priority queue
			sum -= heap.Pop(h2).(int)
		}
		if i <= 2*n {
			suf[i] = sum
		}
	}

	ans := int(^uint(0) >> 1)
	for i := n - 1; i < 2*n; i++ {
		diff := pref[i] - suf[i+1]
		if diff < ans {
			ans = diff
		}
	}
	return int64(ans)
}

func MinimumDifferenceInSumsAfterRemovalOfElements() any {
	return minimumDifference([]int{3, 1, 2})
}
```
