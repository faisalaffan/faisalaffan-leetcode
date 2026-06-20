# 1882 — Process Tasks Using Servers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func AssignTasks(servers []int, tasks []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Heap

**Waktu:** O((m + n) log n) where m = len(tasks), n = len(servers)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Heap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1882: Process Tasks Using Servers
// https://leetcode.com/problems/process-tasks-using-servers/
// Difficulty: Medium

import (
	"container/heap"
	"fmt"
)

type Server struct {
	index int
	weight int
	freeTime int
}

type MinHeapAvailable []Server
func (h MinHeapAvailable) Len() int { return len(h) }
func (h MinHeapAvailable) Less(i, j int) bool {
	if h[i].weight != h[j].weight {
		return h[i].weight < h[j].weight
	}
	return h[i].index < h[j].index
}
func (h MinHeapAvailable) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeapAvailable) Push(x interface{}) { *h = append(*h, x.(Server)) }
func (h *MinHeapAvailable) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MinHeapBusy []Server
func (h MinHeapBusy) Len() int { return len(h) }
func (h MinHeapBusy) Less(i, j int) bool {
	if h[i].freeTime != h[j].freeTime {
		return h[i].freeTime < h[j].freeTime
	}
	if h[i].weight != h[j].weight {
		return h[i].weight < h[j].weight
	}
	return h[i].index < h[j].index
}
func (h MinHeapBusy) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeapBusy) Push(x interface{}) { *h = append(*h, x.(Server)) }
func (h *MinHeapBusy) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	fmt.Println(AssignTasks([]int{3, 3, 2}, []int{1, 2, 3, 2, 1, 2}))
	fmt.Println(AssignTasks([]int{5, 1, 4, 3, 2}, []int{2, 1, 2, 4, 5, 2, 1}))
}

// Time: O((m + n) log n) where m = len(tasks), n = len(servers)
// Space: O(n)
func AssignTasks(servers []int, tasks []int) []int {
	m := len(tasks)

	available := &MinHeapAvailable{}
	heap.Init(available)
	for i, w := range servers {
  // Push ke priority queue
		heap.Push(available, Server{index: i, weight: w, freeTime: 0})
	}

	busy := &MinHeapBusy{}
	heap.Init(busy)

  // Alokasi slice
	result := make([]int, m)
	time := 0

	for j := 0; j < m; j++ {
		time = max(time, j)
		// Release completed servers
		for busy.Len() > 0 && (*busy)[0].freeTime <= time {
  // Pop dari priority queue
			s := heap.Pop(busy).(Server)
  // Push ke priority queue
			heap.Push(available, s)
		}

		if available.Len() == 0 {
			// Jump to next available server's free time
			time = (*busy)[0].freeTime
			for busy.Len() > 0 && (*busy)[0].freeTime <= time {
  // Pop dari priority queue
				s := heap.Pop(busy).(Server)
  // Push ke priority queue
				heap.Push(available, s)
			}
		}

  // Pop dari priority queue
		s := heap.Pop(available).(Server)
		result[j] = s.index
		s.freeTime = time + tasks[j]
  // Push ke priority queue
		heap.Push(busy, s)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
