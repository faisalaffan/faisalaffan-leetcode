# 0743 — Network Delay Time

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func networkDelayTime(times [][]int, n int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Heap / Priority Queue, Stack

**Kompleksitas Waktu:** O(n + E log V)  
**Kompleksitas Ruang:** O(n + E)

> **Untuk fresh graduate:** Kuasai dulu teknik **Heap / Priority Queue** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #743: Network Delay Time
// https://leetcode.com/problems/network-delay-time/
// Difficulty: Medium
// Time: O(n + E log V)
// Space: O(n + E)

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(networkDelayTime([][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2))
	fmt.Println(networkDelayTime([][]int{{1, 2, 1}}, 2, 1))
}

type Edge struct {
	node int
	time int
}

type MinHeap2 []Edge

func (h MinHeap2) Len() int           { return len(h) }
func (h MinHeap2) Less(i, j int) bool { return h[i].time < h[j].time }
func (h MinHeap2) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap2) Push(x interface{}) { *h = append(*h, x.(Edge)) }
func (h *MinHeap2) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func networkDelayTime(times [][]int, n int, k int) int {
  // Membuat matriks/slice 2D untuk DP
	graph := make([][]Edge, n+1)
	for _, t := range times {
		graph[t[0]] = append(graph[t[0]], Edge{t[1], t[2]})
	}

  // Alokasi slice integer
	dist := make([]int, n+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[k] = 0

	h := &MinHeap2{}
	heap.Init(h)
  // Masukkan elemen ke priority queue
	heap.Push(h, Edge{k, 0})

	for h.Len() > 0 {
  // Ambil elemen terkecil/terbesar dari heap
		cur := heap.Pop(h).(Edge)
		if cur.time > dist[cur.node] {
			continue
		}
		for _, e := range graph[cur.node] {
			if nd := cur.time + e.time; nd < dist[e.node] {
				dist[e.node] = nd
  // Masukkan elemen ke priority queue
				heap.Push(h, Edge{e.node, nd})
			}
		}
	}

	maxTime := 0
	for i := 1; i <= n; i++ {
		if dist[i] == math.MaxInt32 {
			return -1
		}
		if dist[i] > maxTime {
			maxTime = dist[i]
		}
	}
	return maxTime
}
```
