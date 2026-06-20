# 3266 — Final Array State After K Multiplication Operations Ii

## Deskripsi

**Soal:** [3266. Final Array State After K Multiplication Operations Ii](https://leetcode.com/problems/final-array-state-after-k-multiplication-operations-ii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** O(n log n * log_k), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Heap (priority queue)

> **Ide Kunci:** //   1. Phase 1: Simulate with a min-heap until min*nums >= max (or k exhausted).

## Solusi Go

```go
package main

// LeetCode #3266: Final Array State After K Multiplication Operations II
// https://leetcode.com/problems/final-array-state-after-k-multiplication-operations-ii/
// Difficulty: Hard
//
// Given an array nums, for each of k operations: find min element (if tie,
// pick first occurrence) and multiply it by multiplier. Apply modulo 1e9+7
// at the end. k can be up to 1e9, so we cannot simulate all operations.
//
// Approach:
//   1. Phase 1: Simulate with a min-heap until min*nums >= max (or k exhausted).
//   2. Phase 2: Remaining ops cycle through all elements. Each element gets
//      base = k/n multiplications; first k%n get one extra.
//   3. Apply fast exponentiation for each element's remaining multiplications.
//
// Time: O(n log n * log_k), Space: O(n)

import (
	"container/heap"
	"fmt"
)

func main() {
	// Example 1: nums=[2,1,3,5,6], k=5, mult=2 => [8,4,6,5,6]
	fmt.Println(getFinalState([]int{2, 1, 3, 5, 6}, 5, 2))
	// Example 2: nums=[1,2], k=3, mult=4 => [16,8]
	fmt.Println(getFinalState([]int{1, 2}, 3, 4))
	// Example 3: multiplier=1 (no change)
	fmt.Println(getFinalState([]int{1, 2, 3}, 100, 1))
	// Example 4: single element
	fmt.Println(getFinalState([]int{5}, 10, 3))
	// Example 5: all same value
	fmt.Println(getFinalState([]int{2, 2, 2}, 3, 2))
}

const MOD3266 = 1_000_000_007

type pair struct {
	val int64
	idx int
}

type minHeap []pair

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].val < h[j].val || (h[i].val == h[j].val && h[i].idx < h[j].idx) }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(pair)) }
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func powMod(x int64, n int64) int64 {
	res := int64(1)
	x %= MOD3266
	for n > 0 {
		if n&1 == 1 {
			res = (res * x) % MOD3266
		}
		x = (x * x) % MOD3266
		n >>= 1
	}
	return res
}

func getFinalState(nums []int, k int, multiplier int) []int64 {
	n := len(nums)
	if multiplier == 1 {
  // Membuat slice untuk menyimpan hasil
		res := make([]int64, n)
		for i, v := range nums {
			res[i] = int64(v) % MOD3266
		}
		return res
	}

	// Phase 1: simulate with heap until min >= max or k exhausted
	maxVal := int64(0)
	h := &minHeap{}
	heap.Init(h)

	for i, v := range nums {
		val := int64(v)
		if val > maxVal {
			maxVal = val
		}
		heap.Push(h, pair{val, i})
	}

	for k > 0 && (*h)[0].val < maxVal {
		p := heap.Pop(h).(pair)
		p.val *= int64(multiplier)
		if p.val > maxVal {
			maxVal = p.val
		}
		heap.Push(h, p)
		k--
	}

	// Phase 2: distribute remaining operations
	// Sort by (value, index)
  // Membuat slice untuk menyimpan hasil
	arr := make([]pair, n)
	for i := 0; i < n; i++ {
		arr[i] = heap.Pop(h).(pair)
	}

	// This wasn't a proper stable sort - we need to sort manually
	// Actually, heap.Pop gives sorted order by value, then index
	// Let's just collect from heap which is already ordered

  // Membuat slice untuk menyimpan hasil
	res := make([]int64, n)
	base := int64(k) / int64(n)
	extra := int64(k) % int64(n)

	for i := 0; i < n; i++ {
		p := arr[i]
		cnt := base
		if int64(i) < extra {
			cnt++
		}
		pow := powMod(int64(multiplier), cnt)
		res[p.idx] = (p.val % MOD3266) * pow % MOD3266
	}

	return res
}
```
