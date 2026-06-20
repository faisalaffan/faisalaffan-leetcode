# 1942 — The Number Of The Smallest Unoccupied Chair

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func SmallestChair(times [][]int, targetFriend int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Heap, Sorting

**Waktu:** O(n log n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1942: The Number of the Smallest Unoccupied Chair
// https://leetcode.com/problems/the-number-of-the-smallest-unoccupied-chair/
// Difficulty: Medium

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SmallestChair([][]int{{1, 4}, {2, 3}, {4, 6}}, 1))
	fmt.Println(SmallestChair([][]int{{3, 10}, {1, 5}, {2, 6}}, 0))
}

type ChairEvent struct {
	time     int
	isLeave  bool
	friend   int
	chairIdx int
}

type MinHeapInt []int

func (h MinHeapInt) Len() int           { return len(h) }
func (h MinHeapInt) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeapInt) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeapInt) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeapInt) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// Time: O(n log n), Space: O(n)
func SmallestChair(times [][]int, targetFriend int) int {
	n := len(times)

	// Pair friend index with their times
	type friend struct {
		arrival    int
		leaving    int
		friendIdx  int
	}
	friends := make([]friend, n)
	for i, t := range times {
		friends[i] = friend{t[0], t[1], i}
	}
  // Custom sort
	sort.Slice(friends, func(i, j int) bool {
		return friends[i].arrival < friends[j].arrival
	})

	// Min heap of available chairs
	available := &MinHeapInt{}
	heap.Init(available)
	for i := 0; i < n; i++ {
  // Push ke priority queue
		heap.Push(available, i)
	}

	// Min heap of occupied chairs (sorted by leave time)
	type occupied struct {
		leaveTime int
		chair     int
	}
	occupiedHeap := make([]occupied, 0)

  // HashMap: O(1) lookup
	chairOf := make(map[int]int) // friend -> chair

	for _, f := range friends {
		// Release chairs of friends who have left
		for len(occupiedHeap) > 0 && occupiedHeap[0].leaveTime <= f.arrival {
			o := occupiedHeap[0]
			occupiedHeap = occupiedHeap[1:]
  // Push ke priority queue
			heap.Push(available, o.chair)
		}

		// Assign smallest available chair
  // Pop dari priority queue
		chair := heap.Pop(available).(int)
		chairOf[f.friendIdx] = chair

		// Insert into occupied (sorted by leave time)
		occupiedHeap = append(occupiedHeap, occupied{f.leaving, chair})
  // Custom sort
		sort.Slice(occupiedHeap, func(i, j int) bool {
			return occupiedHeap[i].leaveTime < occupiedHeap[j].leaveTime
		})
	}

	return chairOf[targetFriend]
}
```
