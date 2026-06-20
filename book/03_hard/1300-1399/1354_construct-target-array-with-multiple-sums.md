# 1354 — Construct Target Array With Multiple Sums

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func isPossible(target []int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Heap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Heap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1354: Construct Target Array With Multiple Sums
// https://leetcode.com/problems/construct-target-array-with-multiple-sums/
// Difficulty: Hard

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func isPossible(target []int) bool {
	if len(target) == 1 {
		return target[0] == 1
	}

	h := &MaxHeap{}
	sum := 0
	for _, v := range target {
		sum += v
  // Push ke priority queue
		heap.Push(h, v)
	}

	for {
  // Pop dari priority queue
		maxVal := heap.Pop(h).(int)
		if maxVal == 1 {
			return true
		}
		rest := sum - maxVal
		if rest == 0 || maxVal <= rest {
			return false
		}
		prev := maxVal % rest
		if prev == 0 {
			prev = rest
		}
		sum = rest + prev
  // Push ke priority queue
		heap.Push(h, prev)
	}
}

func main() {
	// Example 1
	fmt.Println(isPossible([]int{9, 3, 5}))
	// Expected: true

	// Example 2
	fmt.Println(isPossible([]int{1, 1, 1, 2}))
	// Expected: false

	// Example 3
	fmt.Println(isPossible([]int{8, 5}))
	// Expected: true
}
```
