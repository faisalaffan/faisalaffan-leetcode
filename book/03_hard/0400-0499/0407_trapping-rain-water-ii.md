# 0407 — Trapping Rain Water Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func trapRainWater(heightMap [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS, Heap / Priority Queue, Stack

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #407: Trapping Rain Water II
// https://leetcode.com/problems/trapping-rain-water-ii/
// Difficulty: Hard
//
// Use a min-heap (priority queue) to BFS from the border inward. Always
// process the lowest border cell. If a neighbor is lower, water accumulates
// (border height - neighbor height) and the neighbor is "raised" to the
// border height before being pushed back.

import (
	"container/heap"
	"fmt"
)

func main() {
	// Example 1: [[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]] -> 4
	fmt.Println(trapRainWater([][]int{
		{1, 4, 3, 1, 3, 2},
		{3, 2, 1, 3, 2, 4},
		{2, 3, 3, 2, 3, 1},
	}))
	// Example 2: [[3,3,3,3,3],[3,2,2,2,3],[3,2,1,2,3],[3,2,2,2,3],[3,3,3,3,3]] -> 10
	fmt.Println(trapRainWater([][]int{
		{3, 3, 3, 3, 3},
		{3, 2, 2, 2, 3},
		{3, 2, 1, 2, 3},
		{3, 2, 2, 2, 3},
		{3, 3, 3, 3, 3},
	}))
	// Edge: single row
	fmt.Println(trapRainWater([][]int{{1, 2, 1}}))
	// Edge: single column
	fmt.Println(trapRainWater([][]int{{1}, {2}, {1}}))
}

type cell struct {
	h, r, c int
}

type minHeap []cell

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].h < h[j].h }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(cell)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func trapRainWater(heightMap [][]int) int {
	if len(heightMap) < 3 || len(heightMap[0]) < 3 {
		return 0
	}
	rows, cols := len(heightMap), len(heightMap[0])

  // Membuat matriks/slice 2D untuk DP
	visited := make([][]bool, rows)
	for r := range visited {
		visited[r] = make([]bool, cols)
	}

	h := &minHeap{}
	heap.Init(h)

	// Push all border cells
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if r == 0 || r == rows-1 || c == 0 || c == cols-1 {
  // Masukkan elemen ke priority queue
				heap.Push(h, cell{heightMap[r][c], r, c})
				visited[r][c] = true
			}
		}
	}

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	total := 0

	for h.Len() > 0 {
  // Ambil elemen terkecil/terbesar dari heap
		cur := heap.Pop(h).(cell)
		for _, d := range dirs {
			nr, nc := cur.r+d[0], cur.c+d[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && !visited[nr][nc] {
				visited[nr][nc] = true
				if heightMap[nr][nc] < cur.h {
					total += cur.h - heightMap[nr][nc]
  // Masukkan elemen ke priority queue
					heap.Push(h, cell{cur.h, nr, nc})
				} else {
  // Masukkan elemen ke priority queue
					heap.Push(h, cell{heightMap[nr][nc], nr, nc})
				}
			}
		}
	}

	return total
}
```
