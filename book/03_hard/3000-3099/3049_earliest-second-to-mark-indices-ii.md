# 3049 — Earliest Second To Mark Indices Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func earliestSecondToMarkIndices(nums []int, changeIndices []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Binary Search, Heap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3049: Earliest Second to Mark Indices II
// https://leetcode.com/problems/earliest-second-to-mark-indices-ii/
// Difficulty: Hard
//
// Approach: Binary search + greedy with min-heap
// Binary search on the answer (earliest second). For a candidate 'last',
// simulate from right to left: track the first occurrence of each index
// in changeIndices[0:last]. When we encounter a first occurrence, we have
// the option to "apply" the decrement operation (which saves nums[idx]-1
// steps but costs 1 operation slot). Use a min-heap to greedily pick which
// indices to apply decrements to, maximizing the savings.

import (
	"container/heap"
	"fmt"
)

type minHeap3049 []int

func (h minHeap3049) Len() int           { return len(h) }
func (h minHeap3049) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap3049) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap3049) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeap3049) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func earliestSecondToMarkIndices(nums []int, changeIndices []int) int {
	n := len(nums)
	m := len(changeIndices)
	total := int64(0)
	for _, v := range nums {
		total += int64(v)
	}
	total += int64(n)

	check := func(last int) bool {
  // Alokasi slice
		first := make([]int, n)
  // Range loop
		for i := range first {
			first[i] = -1
		}
		for i := 0; i < last; i++ {
			idx := changeIndices[i] - 1
			if first[idx] == -1 {
				first[idx] = i
			}
		}
		pq := &minHeap3049{}
		heap.Init(pq)
		ops := 0
		need := total
		for i := last - 1; i >= 0; i-- {
			idx := changeIndices[i] - 1
			if first[idx] != i {
				ops++
				continue
			}
  // Push ke priority queue
			heap.Push(pq, nums[idx])
			need -= int64(nums[idx]) - 1
			if pq.Len() > ops {
				need += int64((*pq)[0]) - 1
  // Pop dari priority queue
				heap.Pop(pq)
				ops++
			}
		}
		return need <= int64(last)
	}

	lo, hi := 0, m+1
	for lo < hi {
		mid := (lo + hi) / 2
		if check(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	if lo > m {
		return -1
	}
	return lo
}

func main() {
	// Example 1
	fmt.Println("Example 1:", earliestSecondToMarkIndices([]int{2, 2, 3}, []int{1, 2, 3, 1, 2, 3, 1, 2, 3}))
	// Expected: 6

	// Example 2
	fmt.Println("Example 2:", earliestSecondToMarkIndices([]int{2, 2, 0}, []int{2, 2, 2, 2, 3, 2, 2, 1}))
	// Expected: -1 or some value

	// Single element
	fmt.Println("Single:", earliestSecondToMarkIndices([]int{1}, []int{1, 1, 1, 1}))
	// Expected: some value >= 1

	// All zeros
	fmt.Println("All zeros:", earliestSecondToMarkIndices([]int{0, 0, 0}, []int{1, 2, 3, 1, 2, 3}))
	// Expected: some value (zeros are already marked)

	// Simple case
	fmt.Println("Simple:", earliestSecondToMarkIndices([]int{1, 1}, []int{1, 2, 2, 1}))
	// Expected: some value
}
```
