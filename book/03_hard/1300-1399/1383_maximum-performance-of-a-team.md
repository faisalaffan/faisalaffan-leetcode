# 1383 — Maximum Performance Of A Team

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maxPerformance(n int, speed []int, efficiency []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Heap, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Heap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1383: Maximum Performance of a Team
// https://leetcode.com/problems/maximum-performance-of-a-team/
// Difficulty: Hard

import (
	"container/heap"
	"fmt"
	"sort"
)

const mod = 1_000_000_007

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	// Combine speed and efficiency
  // Alokasi slice
	engineers := make([][2]int, n)
  // Range loop
	for i := range engineers {
		engineers[i] = [2]int{efficiency[i], speed[i]}
	}

	// Sort by efficiency descending
  // Custom sort
	sort.Slice(engineers, func(i, j int) bool {
		return engineers[i][0] > engineers[j][0]
	})

	h := &MinHeap{}
	speedSum := 0
	maxPerf := 0

	for _, eng := range engineers {
		eff, sp := eng[0], eng[1]

  // Push ke priority queue
		heap.Push(h, sp)
		speedSum += sp

		if h.Len() > k {
  // Pop dari priority queue
			slow := heap.Pop(h).(int)
			speedSum -= slow
		}

		perf := speedSum * eff
		if perf > maxPerf {
			maxPerf = perf
		}
	}

	return maxPerf % mod
}

func main() {
	// Example 1
	n := 6
	speed := []int{2, 10, 3, 1, 5, 8}
	efficiency := []int{5, 4, 3, 9, 7, 2}
	k := 2
	fmt.Println(maxPerformance(n, speed, efficiency, k))
	// Expected: 60

	// Example 2
	n2 := 6
	speed2 := []int{2, 10, 3, 1, 5, 8}
	efficiency2 := []int{5, 4, 3, 9, 7, 2}
	k2 := 3
	fmt.Println(maxPerformance(n2, speed2, efficiency2, k2))
	// Expected: 68

	// Example 3
	n3 := 6
	speed3 := []int{2, 10, 3, 1, 5, 8}
	efficiency3 := []int{5, 4, 3, 9, 7, 2}
	k3 := 4
	fmt.Println(maxPerformance(n3, speed3, efficiency3, k3))
	// Expected: 72
}
```
