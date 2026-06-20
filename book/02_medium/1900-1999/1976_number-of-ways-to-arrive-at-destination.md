# 1976 — Number Of Ways To Arrive At Destination

## Deskripsi

**Soal:** [1976. Number Of Ways To Arrive At Destination](https://leetcode.com/problems/number-of-ways-to-arrive-at-destination/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(E log V), Space: O(V + E)  
**Kompleksitas Ruang:** O(V + E)

**Algoritma:** Queue (antrian FIFO), Heap (priority queue)

## Solusi Go

```go
package main

// LeetCode #1976: Number of Ways to Arrive at Destination
// https://leetcode.com/problems/number-of-ways-to-arrive-at-destination/
// Difficulty: Medium

import (
	"container/heap"
	"fmt"
)

type Edge struct {
	node int
	time int
}
type Item struct {
	node int
	dist int64
}
type PriorityQueue []Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(Item)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

const mod1976 = 1000000007

func main() {
	fmt.Println(CountPaths(7, [][]int{{0, 6, 7}, {0, 1, 2}, {1, 2, 3}, {1, 3, 3}, {6, 3, 3}, {3, 5, 1}, {6, 5, 1}, {2, 5, 1}, {0, 4, 5}, {4, 6, 2}}))
	fmt.Println(CountPaths(2, [][]int{{1, 0, 10}}))
}

// Time: O(E log V), Space: O(V + E)
func CountPaths(n int, roads [][]int) int {
  // Membuat slice 2D untuk DP/tabel
	graph := make([][]Edge, n)
	for _, r := range roads {
		u, v, t := r[0], r[1], r[2]
		graph[u] = append(graph[u], Edge{v, t})
		graph[v] = append(graph[v], Edge{u, t})
	}

  // Membuat slice untuk menyimpan hasil
	dist := make([]int64, n)
  // Membuat slice untuk menyimpan hasil
	ways := make([]int, n)
  // Iterasi seluruh elemen
	for i := range dist {
		dist[i] = 1 << 62
	}
	dist[0] = 0
	ways[0] = 1

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Item{0, 0})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Item)
		if cur.dist > dist[cur.node] {
			continue
		}
		for _, e := range graph[cur.node] {
			newDist := cur.dist + int64(e.time)
			if newDist < dist[e.node] {
				dist[e.node] = newDist
				ways[e.node] = ways[cur.node]
				heap.Push(pq, Item{e.node, newDist})
			} else if newDist == dist[e.node] {
				ways[e.node] = (ways[e.node] + ways[cur.node]) % mod1976
			}
		}
	}
	return ways[n-1]
}
```
