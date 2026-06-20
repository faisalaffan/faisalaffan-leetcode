# 3672 — Sum Of Weighted Modes In Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func sumOfWeightedModesInSubarrays(nums []int, k int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Heap

**Waktu:** O(n log k)  |  **Ruang:** O(k)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3672: Sum of Weighted Modes in Subarrays
// https://leetcode.com/problems/sum-of-weighted-modes-in-subarrays/
// Difficulty: Medium [Paid]
// Time: O(n log k) | Space: O(k)

import (
	"container/heap"
	"fmt"
)

type pair struct {
	freq int
	val  int
	idx  int // unique index to break ties in heap
}

type maxHeap []pair

func (h maxHeap) Len() int      { return len(h) }
func (h maxHeap) Less(i, j int) bool {
	if h[i].freq != h[j].freq {
		return h[i].freq > h[j].freq
	}
	return h[i].val < h[j].val
}
func (h maxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func sumOfWeightedModesInSubarrays(nums []int, k int) int64 {
	n := len(nums)
  // HashMap: O(1) lookup
	cnt := make(map[int]int)
	pq := &maxHeap{}
	heap.Init(pq)

	add := func(val int, idx int) {
		cnt[val]++
  // Push ke priority queue
		heap.Push(pq, pair{freq: cnt[val], val: val, idx: idx})
	}

	remove := func(val int) {
		cnt[val]--
	}

	getMode := func() (int, int) {
		for pq.Len() > 0 {
			p := (*pq)[0]
			if cnt[p.val] == p.freq {
				return p.val, p.freq
			}
  // Pop dari priority queue
			heap.Pop(pq)
		}
		return 0, 0
	}

	var ans int64 = 0
	globalIdx := 0

	for i := 0; i < n; i++ {
		add(nums[i], globalIdx)
		globalIdx++

		if i >= k {
			remove(nums[i-k])
		}

		if i >= k-1 {
			val, freq := getMode()
			ans += int64(val) * int64(freq)
		}
	}

	return ans
}

func main() {
	fmt.Println(sumOfWeightedModesInSubarrays([]int{1, 2, 2, 3}, 3))
	fmt.Println(sumOfWeightedModesInSubarrays([]int{1, 1, 1, 1}, 2))
	fmt.Println(sumOfWeightedModesInSubarrays([]int{1, 2, 3}, 1))
}
```
