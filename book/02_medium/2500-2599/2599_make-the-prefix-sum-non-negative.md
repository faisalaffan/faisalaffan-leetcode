# 2599 — Make The Prefix Sum Non Negative

## Deskripsi

**Soal:** [2599. Make The Prefix Sum Non Negative](https://leetcode.com/problems/make-the-prefix-sum-non-negative/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Heap (priority queue), Prefix Sum (jumlah kumulatif)

**Fungsi Solusi:** `func makePrefSumNonNegative(nums []int) int`

## Solusi Go

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
			heap.Push(h, v)
		}
		for prefix < 0 {
			// Move the most negative seen to the end
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
