# 1439 — Find The Kth Smallest Sum Of A Matrix With Sorted Rows

## Deskripsi

**Soal:** [1439. Find The Kth Smallest Sum Of A Matrix With Sorted Rows](https://leetcode.com/problems/find-the-kth-smallest-sum-of-a-matrix-with-sorted-rows/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Heap (priority queue)

**Fungsi Solusi:** `func kthSmallest(matrix [][]int, k int) int`

## Solusi Go

```go
package main

// LeetCode #1439: Find the Kth Smallest Sum of a Matrix With Sorted Rows
// https://leetcode.com/problems/find-the-kth-smallest-sum-of-a-matrix-with-sorted-rows/
// Difficulty: Hard

import (
	"container/heap"
	"fmt"
)

type Item struct {
	sum   int
	idx   int // index in the row
	ridx  int // row index
}

type MinHeap []Item

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].sum < h[j].sum }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Item)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kthSmallest(matrix [][]int, k int) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])

	// Start with sums of first row
	h := &MinHeap{}
	for j := 0; j < n; j++ {
		heap.Push(h, Item{sum: matrix[0][j], idx: j, ridx: 0})
	}

	// Merge rows progressively
	for r := 1; r < m; r++ {
		next := &MinHeap{}
		// Take k smallest sums from combining current heap with next row
		count := 0
		for h.Len() > 0 && count < k {
			item := heap.Pop(h).(Item)
			for j := 0; j < n; j++ {
				heap.Push(next, Item{sum: item.sum + matrix[r][j], idx: j, ridx: r})
			}
			count++
		}
		// Keep only k smallest for next iteration
		h = &MinHeap{}
		for i := 0; i < k && next.Len() > 0; i++ {
			heap.Push(h, heap.Pop(next).(Item))
		}
	}

	// Result is kth smallest
	var result int
	for i := 0; i < k; i++ {
		result = heap.Pop(h).(Item).sum
	}
	return result
}

func main() {
	// Example: [[1,3,11],[2,4,6]], k=5 -> 7
	fmt.Println(kthSmallest([][]int{{1, 3, 11}, {2, 4, 6}}, 5))
}
```
