# 2714 — Find Shortest Path With K Hops

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func shortestPathWithKHops(n int, edges [][]int, s int, d int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Heap, Dijkstra

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Heap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2714: Find Shortest Path with K Hops
// https://leetcode.com/problems/find-shortest-path-with-k-hops/
// Difficulty: Hard [Paid]
//
// Approach: Dijkstra with state (node, hopsUsed). We can either pay the edge
// weight or use a "free hop" (cost 0) up to k times.

import (
	"container/heap"
	"fmt"
	"math"
)

type state struct {
	node, dist, hops int
}

type minHeap []state

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(state)) }
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func shortestPathWithKHops(n int, edges [][]int, s int, d int, k int) int {
  // Matriks 2D
	adj := make([][][2]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		adj[u] = append(adj[u], [2]int{v, w})
		adj[v] = append(adj[v], [2]int{u, w})
	}

  // Matriks 2D
	dist := make([][]int, n)
  // Range loop
	for i := range dist {
		dist[i] = make([]int, k+1)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}
	dist[s][0] = 0

	h := &minHeap{}
  // Push ke priority queue
	heap.Push(h, state{s, 0, 0})

	for h.Len() > 0 {
  // Pop dari priority queue
		cur := heap.Pop(h).(state)
		u, du, hops := cur.node, cur.dist, cur.hops
		if du > dist[u][hops] {
			continue
		}
		if u == d {
			return du
		}
		for _, edge := range adj[u] {
			v, w := edge[0], edge[1]
			// Pay cost
			if du+w < dist[v][hops] {
				dist[v][hops] = du + w
  // Push ke priority queue
				heap.Push(h, state{v, du + w, hops})
			}
			// Use free hop
			if hops < k && du < dist[v][hops+1] {
				dist[v][hops+1] = du
  // Push ke priority queue
				heap.Push(h, state{v, du, hops + 1})
			}
		}
	}

	ans := math.MaxInt32
	for i := 0; i <= k; i++ {
		if dist[d][i] < ans {
			ans = dist[d][i]
		}
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(shortestPathWithKHops(5, [][]int{{0, 1, 4}, {0, 2, 2}, {2, 3, 2}, {3, 4, 2}}, 0, 4, 1))
	// Example 2
	fmt.Println(shortestPathWithKHops(4, [][]int{{0, 1, 2}, {1, 2, 1}, {2, 3, 3}}, 0, 3, 2))
	// No free hop needed
	fmt.Println(shortestPathWithKHops(3, [][]int{{0, 1, 5}, {1, 2, 5}}, 0, 2, 0))
	// Unreachable
	fmt.Println(shortestPathWithKHops(3, [][]int{{0, 1, 1}}, 0, 2, 1))
	// Single node
	fmt.Println(shortestPathWithKHops(1, [][]int{}, 0, 0, 0))
}
```
