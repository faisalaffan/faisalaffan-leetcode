# 0871 — Minimum Number Of Refueling Stops

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minRefuelStops(target int, startFuel int, stations [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Heap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Heap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #871: Minimum Number of Refueling Stops
// https://leetcode.com/problems/minimum-number-of-refueling-stops/
// Difficulty: Hard
//
// Greedy with max-heap. Drive as far as current fuel allows. When fuel
// runs out before reaching the next station (or target), pick the station
// with the most fuel seen so far (max-heap) to stop at. This minimizes
// the number of stops.

import (
	"container/heap"
	"fmt"
)

// Max-heap implementation for ints.
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	pq := &MaxHeap{}
	heap.Init(pq)

	stops := 0
	currFuel := startFuel
	i := 0
	n := len(stations)

	for currFuel < target {
		// Add all stations reachable from current position.
		for i < n && stations[i][0] <= currFuel {
  // Push ke priority queue
			heap.Push(pq, stations[i][1])
			i++
		}
		if pq.Len() == 0 {
			return -1 // cannot reach target
		}
		// Stop at the station with the most fuel.
  // Pop dari priority queue
		currFuel += heap.Pop(pq).(int)
		stops++
	}
	return stops
}

func main() {
	// Example 1: target=1, startFuel=1, stations=[] -> 0
	fmt.Println("Test 1:", minRefuelStops(1, 1, [][]int{})) // 0

	// Example 2: target=100, startFuel=1, stations=[[10,100]] -> -1
	fmt.Println("Test 2:", minRefuelStops(100, 1, [][]int{{10, 100}})) // -1

	// Example 3: target=100, startFuel=10, stations=[[10,60],[20,30],[30,30],[60,40]] -> 2
	stations3 := [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}
	fmt.Println("Test 3:", minRefuelStops(100, 10, stations3)) // 2

	// Edge: target=100, startFuel=50, stations=[[25,30]] -> -1 (not enough fuel even after stopping)
	fmt.Println("Test 4:", minRefuelStops(100, 50, [][]int{{25, 30}})) // -1

	// Edge: startFuel already >= target
	fmt.Println("Test 5:", minRefuelStops(50, 100, [][]int{})) // 0
}
```
