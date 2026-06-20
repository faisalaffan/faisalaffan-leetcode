# 0218 — The Skyline Problem

## Deskripsi

**Soal:** [0218. The Skyline Problem](https://leetcode.com/problems/the-skyline-problem/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Queue (antrian FIFO), Heap (priority queue)

**Fungsi Solusi:** `func getSkyline(buildings [][]int) [][]int`

## Solusi Go

```go
package main

// LeetCode #218: The Skyline Problem
// https://leetcode.com/problems/the-skyline-problem/
// Difficulty: Hard

import (
	"container/heap"
	"fmt"
	"sort"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func getSkyline(buildings [][]int) [][]int {
	n := len(buildings)
  // Membuat slice untuk menyimpan hasil
	points := make([][2]int, 0, 2*n)
	for _, b := range buildings {
		points = append(points, [2]int{b[0], -b[2]})
		points = append(points, [2]int{b[1], b[2]})
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i][0] != points[j][0] {
			return points[i][0] < points[j][0]
		}
		return points[i][1] < points[j][1]
	})

  // Membuat slice 2D untuk DP/tabel
	result := make([][]int, 0)
	h := &MaxHeap{}
	heap.Push(h, 0)
	prev := 0

	for _, p := range points {
		x, y := p[0], p[1]
		if y < 0 {
			heap.Push(h, -y)
		} else {
  // Membuat slice untuk menyimpan hasil
			toRemove := make([]int, 0)
			temp := &MaxHeap{}
			for h.Len() > 0 {
				top := heap.Pop(h).(int)
				if top == y {
					break
				}
				toRemove = append(toRemove, top)
			}
			for _, v := range toRemove {
				heap.Push(h, v)
			}
			for temp.Len() > 0 {
				heap.Push(h, heap.Pop(temp).(int))
			}
		}

		// rebuild heap to remove stale heights
		cleaned := &MaxHeap{}
		for h.Len() > 0 {
			top := heap.Pop(h).(int)
			if top != y {
				heap.Push(cleaned, top)
			} else {
				break
			}
		}
		for cleaned.Len() > 0 {
			heap.Push(h, heap.Pop(cleaned).(int))
		}

		if h.Len() > 0 {
			cur := (*h)[0]
			if cur != prev {
				result = append(result, []int{x, cur})
				prev = cur
			}
		}
	}

	return result
}

// Alternative approach using lazy deletion (priority queue)
func getSkyline2(buildings [][]int) [][]int {
	n := len(buildings)
  // Membuat slice untuk menyimpan hasil
	points := make([][2]int, 0, 2*n)
	for _, b := range buildings {
		points = append(points, [2]int{b[0], -b[2]})
		points = append(points, [2]int{b[1], b[2]})
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i][0] != points[j][0] {
			return points[i][0] < points[j][0]
		}
		return points[i][1] < points[j][1]
	})

	// simpler: use a map-based multi-set for heights
  // Membuat map untuk pencarian O(1): key → value
	heights := make(map[int]int)
	heights[0] = 1
	prev := 0
  // Membuat slice 2D untuk DP/tabel
	result := make([][]int, 0)
	// max queue using slice
	maxHeight := func() int {
		max := 0
		for h := range heights {
			if h > max {
				max = h
			}
		}
		return max
	}

	for _, p := range points {
		x, y := p[0], p[1]
		if y < 0 {
			heights[-y]++
		} else {
			heights[y]--
			if heights[y] == 0 {
				delete(heights, y)
			}
		}
		cur := maxHeight()
		if cur != prev {
			result = append(result, []int{x, cur})
			prev = cur
		}
	}
	return result
}

func main() {
	buildings := [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}
	fmt.Println(getSkyline2(buildings))
}
```
