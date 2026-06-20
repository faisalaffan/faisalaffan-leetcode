# 2532 — Time To Cross A Bridge

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findCrossingTime(n int, k int, time [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, BFS, Heap / Priority Queue, Stack

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2532: Time to Cross a Bridge
// https://leetcode.com/problems/time-to-cross-a-bridge/
// Difficulty: Hard
//
// Simulation with 4 priority queues:
//   leftWait: workers on left waiting to cross right (with box)
//   rightWait: workers on right waiting to cross left (without box)
//   leftWork: workers on left picking up boxes
//   rightWork: workers on right putting down boxes
//
// Priority for crossing: right side first, then left side.
// Within same side: smallest (leftToRight+rightToLeft) first, then smallest index.

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	// Example: n=1, k=3, time=[[1,1,2,1],[1,1,3,1],[1,1,4,1]] => 6
	fmt.Println(findCrossingTime(1, 3, [][]int{{1, 1, 2, 1}, {1, 1, 3, 1}, {1, 1, 4, 1}}))
	// Example 2: n=3, k=2, time=[[1,9,1,8],[10,10,10,10]] => 50
	fmt.Println(findCrossingTime(3, 2, [][]int{{1, 9, 1, 8}, {10, 10, 10, 10}}))
	// Edge: single worker, single box
	fmt.Println(findCrossingTime(1, 1, [][]int{{2, 1, 1, 2}}))
	// Edge: no boxes
	fmt.Println(findCrossingTime(0, 2, [][]int{{1, 1, 1, 1}, {2, 2, 2, 2}}))
}

// Worker wrapper for priority queues
type Worker struct {
	idx        int
	efficiency int // leftToRight + rightToLeft (lower = more efficient)
}

type WorkerHeap []Worker

func (h WorkerHeap) Len() int      { return len(h) }
func (h WorkerHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h WorkerHeap) Less(i, j int) bool {
	if h[i].efficiency != h[j].efficiency {
		return h[i].efficiency < h[j].efficiency
	}
	return h[i].idx < h[j].idx
}
func (h *WorkerHeap) Push(x interface{}) { *h = append(*h, x.(Worker)) }
func (h *WorkerHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// TimeHeap for work queues (min-heap by readyTime)
type TimeEvent struct {
	readyTime int
	idx       int
}
type TimeHeap []TimeEvent

func (h TimeHeap) Len() int            { return len(h) }
func (h TimeHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h TimeHeap) Less(i, j int) bool  { return h[i].readyTime < h[j].readyTime }
func (h *TimeHeap) Push(x interface{}) { *h = append(*h, x.(TimeEvent)) }
func (h *TimeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findCrossingTime(n int, k int, time [][]int) int {
  // Edge case: input kosong — langsung return
	if n == 0 {
		return 0
	}

	// Precompute efficiency
  // Alokasi slice integer
	efficiency := make([]int, k)
	for i := 0; i < k; i++ {
		efficiency[i] = time[i][0] + time[i][2]
	}

	// Initial state: all workers on left, start picking up
	leftWork := &TimeHeap{}
	rightWork := &TimeHeap{}
	leftWait := &WorkerHeap{}
	rightWait := &WorkerHeap{}
	heap.Init(leftWork)
	heap.Init(rightWork)
	heap.Init(leftWait)
	heap.Init(rightWait)

	// Initially, workers need to pick up a box before crossing
	for i := 0; i < k; i++ {
  // Masukkan elemen ke priority queue
		heap.Push(leftWork, TimeEvent{readyTime: time[i][1], idx: i}) // pickLeft
	}

	remainingBoxes := n
	currentTime := 0

	// Continue until all boxes are delivered and all workers are on the left
	for remainingBoxes > 0 || rightWork.Len() > 0 || rightWait.Len() > 0 {
		// Flush workers whose work is done into wait queues
		for leftWork.Len() > 0 && (*leftWork)[0].readyTime <= currentTime {
  // Ambil elemen terkecil/terbesar dari heap
			ev := heap.Pop(leftWork).(TimeEvent)
			w := Worker{idx: ev.idx, efficiency: efficiency[ev.idx]}
  // Masukkan elemen ke priority queue
			heap.Push(leftWait, w)
		}
		for rightWork.Len() > 0 && (*rightWork)[0].readyTime <= currentTime {
  // Ambil elemen terkecil/terbesar dari heap
			ev := heap.Pop(rightWork).(TimeEvent)
			w := Worker{idx: ev.idx, efficiency: efficiency[ev.idx]}
  // Masukkan elemen ke priority queue
			heap.Push(rightWait, w)
		}

		if rightWait.Len() > 0 {
			// Worker on right crosses back to left
  // Ambil elemen terkecil/terbesar dari heap
			w := heap.Pop(rightWait).(Worker)
			currentTime += time[w.idx][2] // rightToLeft
			// Worker is now on left, starts picking up
  // Masukkan elemen ke priority queue
			heap.Push(leftWork, TimeEvent{
				readyTime: currentTime + time[w.idx][1], // + pickLeft
				idx:       w.idx,
			})
		} else if leftWait.Len() > 0 && remainingBoxes > 0 {
			// Worker on left crosses to right (with a box, only if boxes remain)
  // Ambil elemen terkecil/terbesar dari heap
			w := heap.Pop(leftWait).(Worker)
			currentTime += time[w.idx][0] // leftToRight
			remainingBoxes--
			// Worker is now on right, starts putting down
  // Masukkan elemen ke priority queue
			heap.Push(rightWork, TimeEvent{
				readyTime: currentTime + time[w.idx][3], // + pickRight
				idx:       w.idx,
			})
		} else {
			// No one waiting, advance time to next work completion
			nextTime := math.MaxInt64
			if leftWork.Len() > 0 && (*leftWork)[0].readyTime < nextTime {
				nextTime = (*leftWork)[0].readyTime
			}
			if rightWork.Len() > 0 && (*rightWork)[0].readyTime < nextTime {
				nextTime = (*rightWork)[0].readyTime
			}
			currentTime = nextTime
		}
	}

	return currentTime
}
```
