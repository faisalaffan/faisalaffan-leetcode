# 1845 — Seat Reservation Manager

## Deskripsi

**Soal:** [1845. Seat Reservation Manager](https://leetcode.com/problems/seat-reservation-manager/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(log n) per operation, Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Heap (priority queue)

**Fungsi Solusi:** `func Constructor(n int) SeatManager`

## Solusi Go

```go
package main

// LeetCode #1845: Seat Reservation Manager
// https://leetcode.com/problems/seat-reservation-manager/
// Difficulty: Medium
// Time: O(log n) per operation, Space: O(n)

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type SeatManager struct {
	h        *MinHeap
	nextSeat int
	n        int
}

func Constructor(n int) SeatManager {
	return SeatManager{
		h:        &MinHeap{},
		nextSeat: 1,
		n:        n,
	}
}

func (sm *SeatManager) Reserve() int {
	if sm.h.Len() > 0 {
		return heap.Pop(sm.h).(int)
	}
	seat := sm.nextSeat
	sm.nextSeat++
	return seat
}

func (sm *SeatManager) Unreserve(seatNumber int) {
	heap.Push(sm.h, seatNumber)
}

func main() {
	sm := Constructor(5)
	fmt.Println(sm.Reserve())    // Expected: 1
	fmt.Println(sm.Reserve())    // Expected: 2
	sm.Unreserve(1)
	fmt.Println(sm.Reserve())    // Expected: 1
	fmt.Println(sm.Reserve())    // Expected: 3
	fmt.Println(sm.Reserve())    // Expected: 4
	sm.Unreserve(2)
	sm.Unreserve(3)
	fmt.Println(sm.Reserve())    // Expected: 2 (smallest unreserved)
}
```
