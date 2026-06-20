# 1845 — Seat Reservation Manager

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(n int) SeatManager
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Heap / Priority Queue, Stack

**Kompleksitas Waktu:** O(log n) per operation, Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Heap / Priority Queue** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

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
  // Ambil elemen terkecil/terbesar dari heap
		return heap.Pop(sm.h).(int)
	}
	seat := sm.nextSeat
	sm.nextSeat++
	return seat
}

func (sm *SeatManager) Unreserve(seatNumber int) {
  // Masukkan elemen ke priority queue
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
