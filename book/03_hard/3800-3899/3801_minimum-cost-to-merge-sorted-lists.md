# 3801 — Minimum Cost To Merge Sorted Lists

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minCost(lists [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** BFS, Heap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **BFS** — sering muncul di interview!

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
  // Push ke priority queue
		heap.Push(h, len(lst))
	}

	total := 0
	for h.Len() > 1 {
  // Pop dari priority queue
		a := heap.Pop(h).(int)
  // Pop dari priority queue
		b := heap.Pop(h).(int)
		cost := a + b
		total += cost
  // Push ke priority queue
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
