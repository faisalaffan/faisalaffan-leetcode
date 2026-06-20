# 2045 — Second Minimum Time To Reach Destination

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func secondMinimum(n int, edges [][]int, time int, change int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Heap, Dijkstra

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Heap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2045: Second Minimum Time to Reach Destination
// https://leetcode.com/problems/second-minimum-time-to-reach-destination/
// Difficulty: Hard
// Approach: Modified Dijkstra / BFS with two distances

import (
	"container/heap"
	"fmt"
	"math"
)

func secondMinimum(n int, edges [][]int, time int, change int) int {
	// Build adjacency list
  // Matriks 2D
	adj := make([][]int, n+1)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// dist1[i] = shortest time to reach i, dist2[i] = second shortest
  // Alokasi slice
	dist1 := make([]int, n+1)
  // Alokasi slice
	dist2 := make([]int, n+1)
  // Range loop
	for i := range dist1 {
		dist1[i] = math.MaxInt32
		dist2[i] = math.MaxInt32
	}

	dist1[1] = 0

	// Min-heap: (time, node)
	pq := &minHeap{}
	heap.Init(pq)
  // Push ke priority queue
	heap.Push(pq, [2]int{0, 1})

	for pq.Len() > 0 {
  // Pop dari priority queue
		cur := heap.Pop(pq).([2]int)
		t := cur[0]
		u := cur[1]

		// If this is stale (worse than dist2), skip
		if t > dist2[u] {
			continue
		}

		// Compute wait time for traffic signal
		// Signal is green in [0, change), red in [change, 2*change), green in [2*change, 3*change), ...
		wait := 0
		cycle := t / change
		if cycle%2 == 1 {
			// Red light, wait until next green
			wait = change - t%change
		}
		nextTime := t + wait + time

		for _, v := range adj[u] {
			if nextTime < dist1[v] {
				dist2[v] = dist1[v]
				dist1[v] = nextTime
  // Push ke priority queue
				heap.Push(pq, [2]int{nextTime, v})
			} else if nextTime > dist1[v] && nextTime < dist2[v] {
				dist2[v] = nextTime
  // Push ke priority queue
				heap.Push(pq, [2]int{nextTime, v})
			}
		}
	}

	return dist2[n]
}

type minHeap [][2]int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	fmt.Println("2045. Second Minimum Time to Reach Destination")

	// Example 1
	n1 := 5
	edges1 := [][]int{{1, 2}, {1, 3}, {1, 4}, {3, 4}, {4, 5}}
	time1 := 3
	change1 := 5
	fmt.Printf("n=%d edges=%v time=%d change=%d → %d (expected 13)\n",
		n1, edges1, time1, change1, secondMinimum(n1, edges1, time1, change1))

	// Example 2
	n2 := 2
	edges2 := [][]int{{1, 2}}
	time2 := 3
	change2 := 2
	fmt.Printf("n=%d edges=%v time=%d change=%d → %d (expected 11)\n",
		n2, edges2, time2, change2, secondMinimum(n2, edges2, time2, change2))
}
```
