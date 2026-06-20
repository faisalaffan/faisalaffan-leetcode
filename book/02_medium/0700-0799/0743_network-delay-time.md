# 0743 — Network Delay Time

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func networkDelayTime(times [][]int, n int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Heap

**Waktu:** O(n + E log V)  |  **Ruang:** O(n + E)

> 🎓 **Fresh Grad Tips:** Kuasai **Heap** — sering muncul di interview!

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
  // Matriks 2D
	graph := make([][]Edge, n+1)
	for _, t := range times {
		graph[t[0]] = append(graph[t[0]], Edge{t[1], t[2]})
	}

  // Alokasi slice
	dist := make([]int, n+1)
  // Range loop
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[k] = 0

	h := &MinHeap2{}
	heap.Init(h)
  // Push ke priority queue
	heap.Push(h, Edge{k, 0})

	for h.Len() > 0 {
  // Pop dari priority queue
		cur := heap.Pop(h).(Edge)
		if cur.time > dist[cur.node] {
			continue
		}
		for _, e := range graph[cur.node] {
			if nd := cur.time + e.time; nd < dist[e.node] {
				dist[e.node] = nd
  // Push ke priority queue
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
