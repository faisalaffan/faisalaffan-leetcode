# 0505 — The Maze Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func TheMazeIi(maze [][]int, start []int, destination []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS, Heap / Priority Queue, Stack, Dijkstra

**Kompleksitas Waktu:** O(m * n * log(m*n)) with Dijkstra, or O(m * n * max(m,n)) with BFS-like  
**Kompleksitas Ruang:** O(m * n)

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #505: The Maze II
// https://leetcode.com/problems/the-maze-ii/
// Difficulty: Medium [Paid]
// Time: O(m * n * log(m*n)) with Dijkstra, or O(m * n * max(m,n)) with BFS-like
// Space: O(m * n)

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	maze := [][]int{
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1},
		{0, 0, 0, 0, 0},
	}
	fmt.Println(TheMazeIi(maze, []int{0, 4}, []int{4, 4}))
	fmt.Println(TheMazeIi(maze, []int{0, 4}, []int{3, 2}))
}

type Item struct {
	r, c, dist int
	index      int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i]; pq[i].index = i; pq[j].index = j }
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func TheMazeIi(maze [][]int, start []int, destination []int) int {
	m, n := len(maze), len(maze[0])
  // Membuat matriks/slice 2D untuk DP
	dist := make([][]int, m)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}

	dist[start[0]][start[1]] = 0
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	pq := &PriorityQueue{}
	heap.Init(pq)
  // Masukkan elemen ke priority queue
	heap.Push(pq, &Item{r: start[0], c: start[1], dist: 0})

	for pq.Len() > 0 {
  // Ambil elemen terkecil/terbesar dari heap
		cur := heap.Pop(pq).(*Item)
		if cur.r == destination[0] && cur.c == destination[1] {
			return cur.dist
		}
		if cur.dist > dist[cur.r][cur.c] {
			continue
		}

		for _, d := range dirs {
			r, c, steps := cur.r, cur.c, 0
			for r+d[0] >= 0 && r+d[0] < m && c+d[1] >= 0 && c+d[1] < n && maze[r+d[0]][c+d[1]] == 0 {
				r += d[0]
				c += d[1]
				steps++
			}
			newDist := cur.dist + steps
			if newDist < dist[r][c] {
				dist[r][c] = newDist
  // Masukkan elemen ke priority queue
				heap.Push(pq, &Item{r: r, c: c, dist: newDist})
			}
		}
	}

	return -1
}
```
