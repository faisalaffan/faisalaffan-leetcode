# 3342 — Find Minimum Time To Reach Last Room Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func minTimeToReach(moveTime [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Heap / Priority Queue, Stack

**Kompleksitas Waktu:** O(m*n*log(m*n)) Space: O(m*n)  
**Kompleksitas Ruang:** O(m*n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Heap / Priority Queue** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3342: Find Minimum Time to Reach Last Room II
// https://leetcode.com/problems/find-minimum-time-to-reach-last-room-ii/
// Difficulty: Medium
// Time: O(m*n*log(m*n)) Space: O(m*n)

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(minTimeToReach([][]int{{0, 4}, {4, 4}})) // 8
}

type State2 struct {
	time int
	r    int
	c    int
}

type MinHeap2 []State2

func (h MinHeap2) Len() int           { return len(h) }
func (h MinHeap2) Less(i, j int) bool { return h[i].time < h[j].time }
func (h MinHeap2) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap2) Push(x interface{}) { *h = append(*h, x.(State2)) }
func (h *MinHeap2) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minTimeToReach(moveTime [][]int) int {
	m, n := len(moveTime), len(moveTime[0])
  // Membuat matriks/slice 2D untuk DP
	dist := make([][]int, m)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = 1 << 60
		}
	}
	dist[0][0] = 0

	h := &MinHeap2{}
	heap.Init(h)
  // Masukkan elemen ke priority queue
	heap.Push(h, State2{0, 0, 0})
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for h.Len() > 0 {
  // Ambil elemen terkecil/terbesar dari heap
		cur := heap.Pop(h).(State2)
		if cur.time > dist[cur.r][cur.c] {
			continue
		}
		if cur.r == m-1 && cur.c == n-1 {
			return cur.time
		}
		for _, d := range dirs {
			nr, nc := cur.r+d[0], cur.c+d[1]
			if nr >= 0 && nr < m && nc >= 0 && nc < n {
				wait := 0
				if cur.time < moveTime[nr][nc] {
					wait = moveTime[nr][nc] - cur.time
				}
				// Move cost: 2 if coming from even cell, 1 if from odd cell
				moveCost := (cur.r+cur.c)%2 + 1
				nt := cur.time + moveCost + wait
				if nt < dist[nr][nc] {
					dist[nr][nc] = nt
  // Masukkan elemen ke priority queue
					heap.Push(h, State2{nt, nr, nc})
				}
			}
		}
	}
	return -1
}
```
