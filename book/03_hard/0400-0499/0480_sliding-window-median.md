# 0480 — Sliding Window Median

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func newLazyHeap(isMax bool) *lazyHeap`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Two Pointer, Sliding Window, Heap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #480: Sliding Window Median
// https://leetcode.com/problems/sliding-window-median/
// Difficulty: Hard
// Approach: Two heaps with lazy deletion (max-heap for left, min-heap for right).
// Maintain balance so left heap has either same count or one more than right heap.

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println("480 - Sliding Window Median")

	// Test cases
	nums1 := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k1 := 3
	fmt.Printf("nums=%v, k=%d -> %v (expected: [1 -1 -1 3 5 6])\n",
		nums1, k1, medianSlidingWindow(nums1, k1))

	nums2 := []int{1, 2, 3, 4, 5}
	k2 := 2
	fmt.Printf("nums=%v, k=%d -> %v (expected: [1.5 2.5 3.5 4.5])\n",
		nums2, k2, medianSlidingWindow(nums2, k2))

	nums3 := []int{1, 2}
	k3 := 1
	fmt.Printf("nums=%v, k=%d -> %v (expected: [1 2])\n",
		nums3, k3, medianSlidingWindow(nums3, k3))

	nums4 := []int{1, 1, 1, 1}
	k4 := 2
	fmt.Printf("nums=%v, k=%d -> %v (expected: [1 1 1])\n",
		nums4, k4, medianSlidingWindow(nums4, k4))

	nums5 := []int{2147483647, 2147483647}
	k5 := 2
	fmt.Printf("nums=%v, k=%d -> %v (expected: [2147483647])\n",
		nums5, k5, medianSlidingWindow(nums5, k5))

	nums6 := []int{5, 2, 3, 1, 4}
	k6 := 3
	fmt.Printf("nums=%v, k=%d -> %v\n",
		nums6, k6, medianSlidingWindow(nums6, k6))
}

// IntHeap for min-heap of int64 (to avoid overflow)
type IntHeap []int64

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int64)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// ReverseIntHeap for max-heap (by nesting IntHeap)
type ReverseIntHeap struct{ IntHeap }

func (h ReverseIntHeap) Less(i, j int) bool { return h.IntHeap[i] > h.IntHeap[j] }

// lazyHeap wraps a heap with a lazy deletion map
type lazyHeap struct {
	h      heap.Interface
	lazy   map[int64]int // value -> count to delete
	size   int
}

func newLazyHeap(isMax bool) *lazyHeap {
	var h heap.Interface
	if isMax {
		h = &ReverseIntHeap{}
	} else {
		h = &IntHeap{}
	}
	heap.Init(h)
	return &lazyHeap{h: h, lazy: make(map[int64]int)}
}

func (lh *lazyHeap) push(val int64) {
  // Push ke priority queue
	heap.Push(lh.h, val)
	lh.size++
}

func (lh *lazyHeap) pop() int64 {
	lh.clean()
	lh.size--
  // Pop dari priority queue
	return heap.Pop(lh.h).(int64)
}

func (lh *lazyHeap) top() int64 {
	lh.clean()
	// We need to peek - let's access the underlying slice
	switch h := lh.h.(type) {
	case *IntHeap:
		return (*h)[0]
	case *ReverseIntHeap:
		return h.IntHeap[0]
	}
	return 0
}

func (lh *lazyHeap) lazyDelete(val int64) {
	lh.lazy[val]++
	lh.size--
}

func (lh *lazyHeap) clean() {
	for {
		var top int64
		switch h := lh.h.(type) {
		case *IntHeap:
			if len(*h) == 0 {
				return
			}
			top = (*h)[0]
		case *ReverseIntHeap:
			if len(h.IntHeap) == 0 {
				return
			}
			top = h.IntHeap[0]
		}

		if count, ok := lh.lazy[top]; ok && count > 0 {
			if count == 1 {
				delete(lh.lazy, top)
			} else {
				lh.lazy[top] = count - 1
			}
  // Pop dari priority queue
			heap.Pop(lh.h)
		} else {
			break
		}
	}
}

func (lh *lazyHeap) len() int {
	return lh.size
}

func medianSlidingWindow(nums []int, k int) []float64 {
	n := len(nums)
	if k == 0 || n == 0 {
		return nil
	}

	result := make([]float64, 0, n-k+1)

	left := newLazyHeap(true)  // max-heap (smaller half)
	right := newLazyHeap(false) // min-heap (larger half)

	// Initialize with first k elements
	for i := 0; i < k; i++ {
		addNum(int64(nums[i]), left, right)
	}

	result = append(result, findMedian(left, right, k))

	for i := k; i < n; i++ {
		// Remove nums[i-k]
		removeNum(int64(nums[i-k]), left, right)
		// Add nums[i]
		addNum(int64(nums[i]), left, right)
		result = append(result, findMedian(left, right, k))
	}

	return result
}

func addNum(val int64, left, right *lazyHeap) {
	if left.len() == 0 || val <= left.top() {
		left.push(val)
	} else {
		right.push(val)
	}

	// Rebalance
	if left.len() > right.len()+1 {
		right.push(left.pop())
	} else if right.len() > left.len() {
		left.push(right.pop())
	}
}

func removeNum(val int64, left, right *lazyHeap) {
	// Determine which heap the value is in and mark for lazy deletion
	if left.len() > 0 && val <= left.top() {
		left.lazyDelete(val)
	} else {
		right.lazyDelete(val)
	}

	// Rebalance
	if left.len() > right.len()+1 {
		right.push(left.pop())
	} else if right.len() > left.len() {
		left.push(right.pop())
	}
}

func findMedian(left, right *lazyHeap, k int) float64 {
	if k%2 == 1 {
		return float64(left.top())
	}
	return float64(left.top()+right.top()) / 2.0
}
```
