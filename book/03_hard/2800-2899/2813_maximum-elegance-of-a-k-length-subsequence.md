# 2813 — Maximum Elegance Of A K Length Subsequence

## Deskripsi

**Soal:** [2813. Maximum Elegance Of A K Length Subsequence](https://leetcode.com/problems/maximum-elegance-of-a-k-length-subsequence/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Heap (priority queue)

**Fungsi Solusi:** `func findMaximumElegance(items [][]int, k int) int64`

## Solusi Go

```go
package main

// LeetCode #2813: Maximum Elegance of a K-Length Subsequence
// https://leetcode.com/problems/maximum-elegance-of-a-k-length-subsequence/
// Difficulty: Hard
//
// Sort items by profit descending. Maintain a min-heap of profits from
// duplicated categories (extra occurrences we can potentially replace).
// For each new item with a new category, try swapping it with the smallest
// profit from a duplicated category to increase distinct category count.
// O(N log N) time, O(N) space.

import (
	"container/heap"
	"fmt"
	"sort"
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

func findMaximumElegance(items [][]int, k int) int64 {
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] > items[j][0]
	})

	h := &MinHeap{}
	heap.Init(h)

	totalProfit := int64(0)
	distinct := 0
  // Membuat map untuk pencarian O(1): key → value
	seenCategory := make(map[int]bool)
	best := int64(0)
	n := len(items)

	for i := 0; i < k && i < n; i++ {
		profit, cat := items[i][0], items[i][1]
		totalProfit += int64(profit)
		if seenCategory[cat] {
			heap.Push(h, profit)
		} else {
			seenCategory[cat] = true
			distinct++
		}
	}

	elegance := totalProfit + int64(distinct)*int64(distinct)
	if elegance > best {
		best = elegance
	}

	for i := k; i < n && h.Len() > 0; i++ {
		profit, cat := items[i][0], items[i][1]
		if seenCategory[cat] {
			continue
		}

		smallest := heap.Pop(h).(int)
		totalProfit -= int64(smallest)
		totalProfit += int64(profit)
		seenCategory[cat] = true
		distinct++

		elegance = totalProfit + int64(distinct)*int64(distinct)
		if elegance > best {
			best = elegance
		}
	}

	return best
}

func main() {
	// Example: items=[[3,2],[5,1],[10,1]], k=2 => 17
	fmt.Println(findMaximumElegance([][]int{{3, 2}, {5, 1}, {10, 1}}, 2))

	// Single category
	fmt.Println(findMaximumElegance([][]int{{1, 1}, {2, 1}, {3, 1}}, 2))

	// All distinct categories
	fmt.Println(findMaximumElegance([][]int{{1, 1}, {2, 2}, {3, 3}}, 2))

	// k = 1
	fmt.Println(findMaximumElegance([][]int{{5, 1}, {10, 2}}, 1))

	// Larger example with duplicate categories
	fmt.Println(findMaximumElegance([][]int{{10, 1}, {10, 1}, {10, 1}, {5, 2}}, 3))

	// k equals number of items
	fmt.Println(findMaximumElegance([][]int{{4, 1}, {3, 2}, {2, 3}}, 3))

	// k = n with duplicates
	fmt.Println(findMaximumElegance([][]int{{5, 1}, {4, 1}, {3, 2}}, 3))

	// All same profit, different categories
	fmt.Println(findMaximumElegance([][]int{{5, 1}, {5, 2}, {5, 3}}, 2))

	// All same profit, same category
	fmt.Println(findMaximumElegance([][]int{{5, 1}, {5, 1}, {5, 1}}, 2))

	// Mix of categories, check swap behavior
	fmt.Println(findMaximumElegance([][]int{
		{100, 1}, {90, 1}, {80, 2}, {70, 3},
	}, 3))

	// Large distinct categories
	fmt.Println(findMaximumElegance([][]int{
		{10, 1}, {9, 2}, {8, 3}, {7, 4}, {6, 5},
	}, 3))
}
```
