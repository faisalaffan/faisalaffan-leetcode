# 3620 — Network Recovery Pathways

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func findMaxPathScore(edges [][]int, online []bool, k int64) int
```

> **💡 Hint:** Binary search on min edge cost + Dijkstra/DP to check feasibility.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Binary Search, Dynamic Programming, Topological Sort, Dijkstra

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3620: Network Recovery Pathways
// https://leetcode.com/problems/network-recovery-pathways/
// Difficulty: Hard
//
// Given a DAG with edges [u,v,cost], boolean array online (which nodes are
// operational), and budget k, find the maximum path score (minimum edge cost
// along the path) from node 0 to node n-1 with total cost <= k.
// Only paths through online intermediate nodes are valid.
//
// Approach: Binary search on min edge cost + Dijkstra/DP to check feasibility.

import "fmt"
import "math"
import "sort"

func main() {
	// Example 1
	fmt.Println(findMaxPathScore([][]int{{0, 1, 3}, {0, 2, 2}, {1, 3, 4}, {2, 3, 1}}, []bool{true, true, true, true}, 6))
	// Example 2
	fmt.Println(findMaxPathScore([][]int{{0, 1, 5}, {1, 2, 3}}, []bool{true, true, true}, 7))
	// Edge: no valid path
	fmt.Println(findMaxPathScore([][]int{{0, 1, 5}}, []bool{true, false}, 10))
	// Edge: single node
	fmt.Println(findMaxPathScore([][]int{}, []bool{true}, 0))
}

func findMaxPathScore(edges [][]int, online []bool, k int64) int {
	n := len(online)
	if n <= 1 {
		return 0
	}

	// Build adjacency
	type edge struct{ to, cost int }
  // Membuat matriks/slice 2D untuk DP
	adj := make([][]edge, n)
  // Alokasi slice integer
	costs := make([]int, 0)
	for _, e := range edges {
		u, v, c := e[0], e[1], e[2]
		adj[u] = append(adj[u], edge{v, c})
		costs = append(costs, c)
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(costs)

	if !online[0] || !online[n-1] {
		return -1
	}

	// Check if path with minEdge >= x and total cost <= k exists
	check := func(x int) bool {
		// DP: dist[node] = min total cost from node 0 to node
  // Alokasi slice integer
		dist := make([]int64, n)
  // Range loop: iterasi dengan indeks + nilai
		for i := range dist {
			dist[i] = math.MaxInt64
		}
		dist[0] = 0

		// Topological DP (graph is DAG)
		// Simple DP for DAG: visit nodes in order
		for u := 0; u < n; u++ {
			if dist[u] == math.MaxInt64 {
				continue
			}
			if !online[u] && u != 0 {
				continue
			}
			for _, e := range adj[u] {
				if !online[e.to] && e.to != n-1 {
					continue
				}
				if e.cost >= x {
					nd := dist[u] + int64(e.cost)
					if nd < dist[e.to] {
						dist[e.to] = nd
					}
				}
			}
		}

		return dist[n-1] <= k
	}

	// Binary search on min edge cost
	left, right := 0, len(costs)-1
	result := -1

	for left <= right {
		mid := (left + right) / 2
		if check(costs[mid]) {
			result = costs[mid]
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if result == -1 {
		// Check if zero threshold path exists
		if check(0) {
			return 0
		}
		return -1
	}
	return result
}
```
