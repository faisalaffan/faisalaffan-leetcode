# 3763 — Maximum Total Sum With Threshold Constraints

## Deskripsi

**Soal:** [3763. Maximum Total Sum With Threshold Constraints](https://leetcode.com/problems/maximum-total-sum-with-threshold-constraints/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Heap (priority queue)

**Fungsi Solusi:** `func maximumTotalSumWithThresholdConstraints(nums []int, threshold []int) int`

## Solusi Go

```go
package main

// LeetCode #3763: Maximum Total Sum with Threshold Constraints
// https://leetcode.com/problems/maximum-total-sum-with-threshold-constraints/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"container/heap"
	"fmt"
	"sort"
)

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maximumTotalSumWithThresholdConstraints(nums []int, threshold []int) int {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	items := make([][2]int, n)
	for i := 0; i < n; i++ {
		items[i] = [2]int{threshold[i], nums[i]}
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})

	h := &maxHeap{}
	heap.Init(h)
	idx := 0
	total := 0

	for step := 1; step <= n; step++ {
		for idx < n && items[idx][0] <= step {
			heap.Push(h, items[idx][1])
			idx++
		}
		if h.Len() == 0 {
			break
		}
		total += heap.Pop(h).(int)
	}
	return total
}

func main() {
	fmt.Println(maximumTotalSumWithThresholdConstraints([]int{5, 3, 4}, []int{3, 1, 2}))
	fmt.Println(maximumTotalSumWithThresholdConstraints([]int{10, 20, 30}, []int{3, 2, 1}))
	fmt.Println(maximumTotalSumWithThresholdConstraints([]int{1, 2}, []int{1, 1}))
}
```
