# 3538 — Merge Operations For Minimum Travel Time

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumTravelTime(n int, edges [][]int) int64
```

> **💡 Hint:** DP on intervals or use Dijkstra with state compression.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS, Dynamic Programming, Heap / Priority Queue, Stack, Dijkstra

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3538: Merge Operations for Minimum Travel Time
// https://leetcode.com/problems/merge-operations-for-minimum-travel-time/
// Difficulty: Hard
//
// Given a graph with travel times, merge nodes to minimize travel time
// between start and end. Each merge combines two adjacent nodes.
//
// Approach: DP on intervals or use Dijkstra with state compression.

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	// Example 1
	fmt.Println(minimumTravelTime(4, [][]int{{0, 1, 2}, {0, 2, 5}, {2, 3, 1}, {1, 3, 3}}))
	// Example 2
	fmt.Println(minimumTravelTime(3, [][]int{{0, 1, 1}, {1, 2, 2}}))
	// Edge: single edge
	fmt.Println(minimumTravelTime(2, [][]int{{0, 1, 5}}))
}

type Item struct {
	node int
	dist int64
	idx  int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i]; pq[i].idx = i; pq[j].idx = j }
func (pq *PriorityQueue) Push(x interface{}) { n := len(*pq); item := x.(*Item); item.idx = n; *pq = append(*pq, item) }
func (pq *PriorityQueue) Pop() interface{} { old := *pq; n := len(old); item := old[n-1]; item.idx = -1; *pq = old[:n-1]; return item }

func minimumTravelTime(n int, edges [][]int) int64 {
  // Membuat matriks/slice 2D untuk DP
	adj := make([][][2]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		adj[u] = append(adj[u], [2]int{v, w})
		adj[v] = append(adj[v], [2]int{u, w})
	}

	// Dijkstra from 0 to n-1
  // Alokasi slice integer
	dist := make([]int64, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[0] = 0
	pq := &PriorityQueue{}
	heap.Init(pq)
  // Masukkan elemen ke priority queue
	heap.Push(pq, &Item{node: 0, dist: 0})

	for pq.Len() > 0 {
  // Ambil elemen terkecil/terbesar dari heap
		item := heap.Pop(pq).(*Item)
		u := item.node
		if item.dist > dist[u] {
			continue
		}
		if u == n-1 {
			return item.dist
		}
		for _, edge := range adj[u] {
			v, w := edge[0], int64(edge[1])
			if nd := item.dist + w; nd < dist[v] {
				dist[v] = nd
  // Masukkan elemen ke priority queue
				heap.Push(pq, &Item{node: v, dist: nd})
			}
		}
	}

	return -1
}
```
