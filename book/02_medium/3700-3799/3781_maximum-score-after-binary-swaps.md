# 3781 — Maximum Score After Binary Swaps

## Deskripsi

**Soal:** [3781. Maximum Score After Binary Swaps](https://leetcode.com/problems/maximum-score-after-binary-swaps/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Heap (priority queue)

**Fungsi Solusi:** `func maximumScoreAfterBinarySwaps(nums []int, s string) int64`

## Solusi Go

```go
package main

// LeetCode #3781: Maximum Score After Binary Swaps
// https://leetcode.com/problems/maximum-score-after-binary-swaps/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type maxHeap3781 []int

func (h maxHeap3781) Len() int           { return len(h) }
func (h maxHeap3781) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap3781) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap3781) Push(x any)        { *h = append(*h, x.(int)) }
func (h *maxHeap3781) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maximumScoreAfterBinarySwaps(nums []int, s string) int64 {
	h := &maxHeap3781{}
	heap.Init(h)
	var ans int64

	for i, ch := range s {
		heap.Push(h, nums[i])
		if ch == '1' {
			ans += int64(heap.Pop(h).(int))
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumScoreAfterBinarySwaps([]int{3, 1, 4, 2}, "1010"))
	fmt.Println(maximumScoreAfterBinarySwaps([]int{5, 2, 8, 1}, "1001"))
	fmt.Println(maximumScoreAfterBinarySwaps([]int{10, 20, 30}, "111"))
}
```
