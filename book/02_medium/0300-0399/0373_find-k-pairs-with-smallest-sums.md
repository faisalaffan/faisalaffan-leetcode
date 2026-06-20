# 0373 — Find K Pairs With Smallest Sums

## Deskripsi

**Soal:** [0373. Find K Pairs With Smallest Sums](https://leetcode.com/problems/find-k-pairs-with-smallest-sums/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(k log min(k, n))  
**Kompleksitas Ruang:** O(k)

**Algoritma:** Heap (priority queue)

**Fungsi Solusi:** `func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int`

## Solusi Go

```go
package main

// LeetCode #373: Find K Pairs with Smallest Sums
// https://leetcode.com/problems/find-k-pairs-with-smallest-sums/
// Difficulty: Medium
// Time: O(k log min(k, n)) | Space: O(k)

import (
	"container/heap"
	"fmt"
)

type pair struct {
	i, j int
	sum  int
}

type minHeap []pair

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].sum < h[j].sum }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
		return [][]int{}
	}

	h := &minHeap{}
	heap.Init(h)

	// Push first element of nums1 paired with each element of nums2
	for j := 0; j < len(nums2) && j < k; j++ {
		heap.Push(h, pair{0, j, nums1[0] + nums2[j]})
	}

  // Membuat slice 2D untuk DP/tabel
	result := make([][]int, 0, k)
	for h.Len() > 0 && len(result) < k {
		p := heap.Pop(h).(pair)
		result = append(result, []int{nums1[p.i], nums2[p.j]})
		if p.i+1 < len(nums1) {
			heap.Push(h, pair{p.i + 1, p.j, nums1[p.i+1] + nums2[p.j]})
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3))
	// Expected: [[1,2],[1,4],[1,6]]

	// Test case 2
	fmt.Println("Test 2:", kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 2))
	// Expected: [[1,1],[1,1]]

	// Test case 3
	fmt.Println("Test 3:", kSmallestPairs([]int{1, 2}, []int{3}, 3))
	// Expected: [[1,3],[2,3]]
}
```
