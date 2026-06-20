# 3092 — Most Frequent Ids

## Deskripsi

**Soal:** [3092. Most Frequent Ids](https://leetcode.com/problems/most-frequent-ids/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Heap (priority queue)

**Fungsi Solusi:** `func mostFrequentIDs(nums []int, freq []int) []int64`

## Solusi Go

```go
package main

// LeetCode #3092: Most Frequent IDs
// https://leetcode.com/problems/most-frequent-ids/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type Item struct {
	id    int
	count int64
	idx   int
}

type MaxHeap []*Item

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].count > h[j].count }
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].idx = i
	h[j].idx = j
}
func (h *MaxHeap) Push(x any) {
	n := len(*h)
	item := x.(*Item)
	item.idx = n
	*h = append(*h, item)
}
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.idx = -1
	*h = old[:n-1]
	return item
}

func mostFrequentIDs(nums []int, freq []int) []int64 {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	ans := make([]int64, n)
  // Membuat map untuk pencarian O(1): key → value
	counts := make(map[int]int64)
	h := &MaxHeap{}
	heap.Init(h)

	for i := 0; i < n; i++ {
		id := nums[i]
		counts[id] += int64(freq[i])
		heap.Push(h, &Item{id: id, count: counts[id]})
		for h.Len() > 0 && (*h)[0].count != counts[(*h)[0].id] {
			heap.Pop(h)
		}
		if h.Len() > 0 {
			ans[i] = (*h)[0].count
		}
	}
	return ans
}

func main() {
	fmt.Println(mostFrequentIDs([]int{1, 2, 3, 2, 2, 1, 3}, []int{-2, 3, -3, 5, 1, -3, 1}))
	fmt.Println(mostFrequentIDs([]int{2, 3, 2, 1}, []int{3, 2, -3, 1}))
}
```
