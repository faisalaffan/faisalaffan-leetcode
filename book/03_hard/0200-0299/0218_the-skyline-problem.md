# 0218 — The Skyline Problem

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func getSkyline(buildings [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, BFS, Heap / Priority Queue, Stack

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

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
  // Alokasi slice integer
	points := make([][2]int, 0, 2*n)
	for _, b := range buildings {
		points = append(points, [2]int{b[0], -b[2]})
		points = append(points, [2]int{b[1], b[2]})
	}

  // Custom sort dengan comparator
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] != points[j][0] {
			return points[i][0] < points[j][0]
		}
		return points[i][1] < points[j][1]
	})

  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, 0)
	h := &MaxHeap{}
  // Masukkan elemen ke priority queue
	heap.Push(h, 0)
	prev := 0

	for _, p := range points {
		x, y := p[0], p[1]
		if y < 0 {
  // Masukkan elemen ke priority queue
			heap.Push(h, -y)
		} else {
  // Alokasi slice integer
			toRemove := make([]int, 0)
			temp := &MaxHeap{}
			for h.Len() > 0 {
  // Ambil elemen terkecil/terbesar dari heap
				top := heap.Pop(h).(int)
				if top == y {
					break
				}
				toRemove = append(toRemove, top)
			}
			for _, v := range toRemove {
  // Masukkan elemen ke priority queue
				heap.Push(h, v)
			}
			for temp.Len() > 0 {
  // Masukkan elemen ke priority queue
				heap.Push(h, heap.Pop(temp).(int))
			}
		}

		// rebuild heap to remove stale heights
		cleaned := &MaxHeap{}
		for h.Len() > 0 {
  // Ambil elemen terkecil/terbesar dari heap
			top := heap.Pop(h).(int)
			if top != y {
  // Masukkan elemen ke priority queue
				heap.Push(cleaned, top)
			} else {
				break
			}
		}
		for cleaned.Len() > 0 {
  // Masukkan elemen ke priority queue
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
  // Alokasi slice integer
	points := make([][2]int, 0, 2*n)
	for _, b := range buildings {
		points = append(points, [2]int{b[0], -b[2]})
		points = append(points, [2]int{b[1], b[2]})
	}

  // Custom sort dengan comparator
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] != points[j][0] {
			return points[i][0] < points[j][0]
		}
		return points[i][1] < points[j][1]
	})

	// simpler: use a map-based multi-set for heights
  // Membuat map (HashMap) — pencarian O(1)
	heights := make(map[int]int)
	heights[0] = 1
	prev := 0
  // Membuat matriks/slice 2D untuk DP
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
