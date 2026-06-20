# 1199 — Minimum Time To Build Blocks

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minimumTimeToBuildBlocks(blocks []int, split int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Heap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Heap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1199: Minimum Time to Build Blocks
// https://leetcode.com/problems/minimum-time-to-build-blocks/
// Difficulty: Hard [Paid]
//
// You have blocks, each requiring a certain number of time units to build.
// You start with one worker. In each unit of time, a worker can either build
// one block (taking the block's time) or split into two workers (taking
// `split` time). Workers work in parallel. Find the minimum time to build
// all blocks.

import (
	"container/heap"
	"fmt"
)

// MinHeap implements heap.Interface for ints (min-heap).
type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)         { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	// Example 1
	fmt.Println(minimumTimeToBuildBlocks([]int{1, 2}, 5)) // 7

	// Example 2
	fmt.Println(minimumTimeToBuildBlocks([]int{1, 2, 3}, 5)) // 12

	// Single block
	fmt.Println(minimumTimeToBuildBlocks([]int{7}, 10)) // 7

	// All equal small
	fmt.Println(minimumTimeToBuildBlocks([]int{1, 1, 1, 1}, 2)) // 5

	// Larger split cost
	fmt.Println(minimumTimeToBuildBlocks([]int{1, 2, 4, 8}, 10)) // 28
}

// minimumTimeToBuildBlocks returns the minimum time to build all blocks.
//
// This is a Huffman coding-like problem. We start with all blocks as separate
// "jobs." At each step, we pick the two smallest jobs and merge them: the
// combined job takes max(a, b) + split time. This represents a worker spending
// `split` time to split, creating two workers that build the two blocks in
// parallel. The + split accounts for the split operation, and the max(a, b)
// accounts for the longer block build time (parallel work).
//
// We use a min-heap to always merge the smallest (fastest) jobs first, which
// minimizes the critical path.
func minimumTimeToBuildBlocks(blocks []int, split int) int {
	h := &MinHeap{}
	heap.Init(h)

	for _, b := range blocks {
  // Push ke priority queue
		heap.Push(h, b)
	}

	for h.Len() > 1 {
  // Pop dari priority queue
		a := heap.Pop(h).(int)
  // Pop dari priority queue
		b := heap.Pop(h).(int)
		combined := max(a, b) + split
  // Push ke priority queue
		heap.Push(h, combined)
	}

  // Pop dari priority queue
	return heap.Pop(h).(int)
}
```
