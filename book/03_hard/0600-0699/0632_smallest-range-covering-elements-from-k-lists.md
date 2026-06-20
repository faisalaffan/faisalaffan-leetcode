# 0632 — Smallest Range Covering Elements From K Lists

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func smallestRange(nums [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Heap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Heap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #632: Smallest Range Covering Elements from K Lists
// https://leetcode.com/problems/smallest-range-covering-elements-from-k-lists/
// Difficulty: Hard

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	// Test cases
	testCases := []struct {
		nums [][]int
		want []int
	}{
		{[][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}}, []int{20, 24}},
		{[][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}, []int{1, 1}},
		{[][]int{{10, 10}, {11, 11}}, []int{10, 11}},
		{[][]int{{1}, {2}, {3}, {4}}, []int{1, 4}},
		{[][]int{{1, 3, 5, 7}, {2, 4, 6, 8}}, []int{1, 2}},
		{[][]int{{1, 4, 7, 10}, {2, 5, 8, 11}, {3, 6, 9, 12}}, []int{1, 3}},
		{[][]int{{1, 2}, {1, 3}, {1, 5}}, []int{1, 1}},
	}

	for _, tc := range testCases {
		got := smallestRange(tc.nums)
		status := "PASS"
		if len(got) == 2 && len(tc.want) == 2 && (got[0] != tc.want[0] || got[1] != tc.want[1]) {
			status = "FAIL"
		}
		fmt.Printf("%s: smallestRange(%v) = %v (want %v)\n", status, tc.nums, got, tc.want)
	}
}

type element struct {
	val   int
	list  int
	index int
}

type minHeap []element

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(element))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func smallestRange(nums [][]int) []int {
	k := len(nums)
	h := &minHeap{}
	heap.Init(h)

	maxVal := math.MinInt32
	for i := 0; i < k; i++ {
  // Push ke priority queue
		heap.Push(h, element{val: nums[i][0], list: i, index: 0})
		if nums[i][0] > maxVal {
			maxVal = nums[i][0]
		}
	}

	start, end := 0, math.MaxInt32

	for {
  // Pop dari priority queue
		minElem := heap.Pop(h).(element)
		curStart := minElem.val
		curEnd := maxVal

		// Update range if smaller
		if curEnd-curStart < end-start {
			start = curStart
			end = curEnd
		}

		// If no more elements in this list, we can't cover all lists
		if minElem.index+1 >= len(nums[minElem.list]) {
			break
		}

		// Push next element from the same list
		nextVal := nums[minElem.list][minElem.index+1]
  // Push ke priority queue
		heap.Push(h, element{val: nextVal, list: minElem.list, index: minElem.index + 1})
		if nextVal > maxVal {
			maxVal = nextVal
		}
	}

	return []int{start, end}
}
```
