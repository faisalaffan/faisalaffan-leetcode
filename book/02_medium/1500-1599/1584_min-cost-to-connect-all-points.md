# 1584 — Min Cost To Connect All Points

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinCostConnectPoints(points [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Heap / Priority Queue, Stack

**Kompleksitas Waktu:** O(N^2 log N), Space: O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **Heap / Priority Queue** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1584: Min Cost to Connect All Points
// https://leetcode.com/problems/min-cost-to-connect-all-points/
// Difficulty: Medium

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(MinCostConnectPoints([][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}))
	fmt.Println(MinCostConnectPoints([][]int{{3, 12}, {-2, 5}, {-4, 1}}))
	fmt.Println(MinCostConnectPoints([][]int{{0, 0}, {1, 1}, {1, 0}, {-1, 1}}))
}

type Edge struct {
	to   int
	dist int
}

type PQItem struct {
	dist  int
	node  int
	index int
}

type PQ []*PQItem

func (pq PQ) Len() int           { return len(pq) }
func (pq PQ) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*PQItem)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[:n-1]
	return item
}

func MinCostConnectPoints(points [][]int) int {
	// Time: O(N^2 log N), Space: O(N)
	// Prim's algorithm
	n := len(points)
	if n <= 1 {
		return 0
	}

	visited := make([]bool, n)
	pq := &PQ{}
	heap.Init(pq)
  // Masukkan elemen ke priority queue
	heap.Push(pq, &PQItem{dist: 0, node: 0})

	totalCost := 0
	edgesUsed := 0

	for pq.Len() > 0 && edgesUsed < n {
  // Ambil elemen terkecil/terbesar dari heap
		item := heap.Pop(pq).(*PQItem)
		if visited[item.node] {
			continue
		}
		visited[item.node] = true
		totalCost += item.dist
		edgesUsed++

		for i := 0; i < n; i++ {
			if !visited[i] {
				dist := abs(points[item.node][0]-points[i][0]) + abs(points[item.node][1]-points[i][1])
  // Masukkan elemen ke priority queue
				heap.Push(pq, &PQItem{dist: dist, node: i})
			}
		}
	}

	return totalCost
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
