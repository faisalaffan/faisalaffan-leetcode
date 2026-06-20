# 2402 — Meeting Rooms Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func mostBooked(n int, meetings [][]int) int
```

> **💡 Hint:** Use two min-heaps:

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Heap / Priority Queue, Stack

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Heap / Priority Queue** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2402: Meeting Rooms III
// https://leetcode.com/problems/meeting-rooms-iii/
// Difficulty: Hard
//
// You have n meeting rooms numbered 0 to n-1. Given meetings[start_i, end_i),
// assign each meeting to the smallest-numbered available room. If no room is
// available, delay the meeting until a room is free (keeping duration same).
// Return the room that hosts the most meetings.
//
// Approach: Use two min-heaps:
//   - available: room indices (min-heap by room number)
//   - busy: (endTime, roomIndex) (min-heap by endTime)
// Process meetings in sorted order by start time. For each meeting, release
// all busy rooms that are now free. Assign the meeting to the smallest
// available room. If no room is free, use the earliest-freed room.

import (
	"container/heap"
	"fmt"
	"sort"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type Busy struct {
	endTime int
	room    int
}

type BusyHeap []Busy

func (h BusyHeap) Len() int           { return len(h) }
func (h BusyHeap) Less(i, j int) bool { return h[i].endTime < h[j].endTime }
func (h BusyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *BusyHeap) Push(x interface{}) {
	*h = append(*h, x.(Busy))
}

func (h *BusyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func mostBooked(n int, meetings [][]int) int {
	// Sort meetings by start time
  // Custom sort dengan comparator
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	avail := &IntHeap{}
	for i := 0; i < n; i++ {
  // Masukkan elemen ke priority queue
		heap.Push(avail, i)
	}

	busy := &BusyHeap{}
  // Alokasi slice integer
	count := make([]int, n)

	for _, m := range meetings {
		start, end := m[0], m[1]
		duration := end - start

		// Release rooms that have finished by now
		for busy.Len() > 0 && (*busy)[0].endTime <= start {
  // Ambil elemen terkecil/terbesar dari heap
			b := heap.Pop(busy).(Busy)
  // Masukkan elemen ke priority queue
			heap.Push(avail, b.room)
		}

		var room int
		if avail.Len() > 0 {
			// Room available, use smallest-numbered
  // Ambil elemen terkecil/terbesar dari heap
			room = heap.Pop(avail).(int)
		} else {
			// No room available, pick the soonest free room and delay
  // Ambil elemen terkecil/terbesar dari heap
			b := heap.Pop(busy).(Busy)
			room = b.room
			// Meeting is delayed; duration stays same
			start = b.endTime
		}

		count[room]++
  // Masukkan elemen ke priority queue
		heap.Push(busy, Busy{endTime: start + duration, room: room})
	}

	// Find room with max count, smallest number in case of tie
	maxRoom, maxCount := 0, count[0]
	for i := 1; i < n; i++ {
		if count[i] > maxCount {
			maxCount = count[i]
			maxRoom = i
		}
	}

	return maxRoom
}

func main() {
	// Example 1
	n1 := 2
	meetings1 := [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}}
	fmt.Println(mostBooked(n1, meetings1))

	// Example 2
	n2 := 3
	meetings2 := [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}}
	fmt.Println(mostBooked(n2, meetings2))

	// Single room
	n3 := 1
	meetings3 := [][]int{{0, 5}, {2, 3}, {4, 7}}
	fmt.Println(mostBooked(n3, meetings3))

	// All meetings overlap
	n4 := 2
	meetings4 := [][]int{{0, 10}, {0, 5}, {0, 7}}
	fmt.Println(mostBooked(n4, meetings4))
}
```
