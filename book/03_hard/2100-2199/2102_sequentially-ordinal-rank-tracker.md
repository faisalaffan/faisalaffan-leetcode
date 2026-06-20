# 2102 — Sequentially Ordinal Rank Tracker

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func Constructor() SORTracker`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Heap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Heap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2102: Sequentially Ordinal Rank Tracker
// https://leetcode.com/problems/sequentially-ordinal-rank-tracker/
// Difficulty: Hard
//
// Two-heap approach:
// - low (min-heap by score ASC, name DESC): contains top-k items
// - high (min-heap by score DESC, name ASC): contains items beyond top-k

import (
	"container/heap"
	"fmt"
)

type location struct {
	name  string
	score int
}

type lowHeap []location

func (h lowHeap) Len() int      { return len(h) }
func (h lowHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h lowHeap) Less(i, j int) bool {
	if h[i].score != h[j].score {
		return h[i].score < h[j].score
	}
	return h[i].name > h[j].name
}
func (h *lowHeap) Push(x any)   { *h = append(*h, x.(location)) }
func (h *lowHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type highHeap []location

func (h highHeap) Len() int      { return len(h) }
func (h highHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h highHeap) Less(i, j int) bool {
	if h[i].score != h[j].score {
		return h[i].score > h[j].score
	}
	return h[i].name < h[j].name
}
func (h *highHeap) Push(x any)   { *h = append(*h, x.(location)) }
func (h *highHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type SORTracker struct {
	low     lowHeap
	high    highHeap
	queries int
}

func Constructor() SORTracker {
	return SORTracker{}
}

func (t *SORTracker) Add(name string, score int) {
  // Push ke priority queue
	heap.Push(&t.low, location{name, score})
	if len(t.low) > t.queries {
  // Push ke priority queue
		heap.Push(&t.high, heap.Pop(&t.low))
	}
}

func (t *SORTracker) Get() string {
	t.queries++
	for len(t.low) < t.queries {
  // Push ke priority queue
		heap.Push(&t.low, heap.Pop(&t.high))
	}
	return t.low[0].name
}

func main() {
	tracker := Constructor()
	tracker.Add("bradford", 2)
	tracker.Add("branford", 3)
	fmt.Println(tracker.Get())
	tracker.Add("alps", 2)
	fmt.Println(tracker.Get())
	tracker.Add("orland", 2)
	fmt.Println(tracker.Get())
	tracker.Add("orlando", 3)
	fmt.Println(tracker.Get())
	tracker.Add("alpine", 2)
	fmt.Println(tracker.Get())
	fmt.Println(tracker.Get())
}
```
