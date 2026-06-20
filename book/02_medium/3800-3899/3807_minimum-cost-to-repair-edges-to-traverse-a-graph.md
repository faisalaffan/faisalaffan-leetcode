# 3807 — Minimum Cost To Repair Edges To Traverse A Graph

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumCostToRepairEdgesToTraverseAGraph(n int, edges [][]int, k int) int
```

> **💡 Hint:** Binary search on cost + BFS to check reachability within k edges.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** O((N+M) * log M)  
**Kompleksitas Ruang:** O(N+M)

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3807: Minimum Cost to Repair Edges to Traverse a Graph
// https://leetcode.com/problems/minimum-cost-to-repair-edges-to-traverse-a-graph/
// Difficulty: Medium [Paid]
// Time: O((N+M) * log M) | Space: O(N+M)
// Approach: Binary search on cost + BFS to check reachability within k edges.

import (
	"container/list"
	"fmt"
	"sort"
)

func MinimumCostToRepairEdgesToTraverseAGraph(n int, edges [][]int, k int) int {
	// Sort edges by repair cost
  // Custom sort dengan comparator
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	m := len(edges)

	// Binary search on edge cost threshold
	lo, hi := 0, m
	ans := -1

	for lo <= hi {
		mid := (lo + hi) / 2
		if mid >= m {
			break // can't use edges beyond available
		}
		costLimit := edges[mid][2]

		// Build graph with edges <= costLimit
  // Membuat matriks/slice 2D untuk DP
		adj := make([][]int, n)
		for _, e := range edges {
			if e[2] <= costLimit {
				u, v := e[0], e[1]
				adj[u] = append(adj[u], v)
				adj[v] = append(adj[v], u)
			}
		}

		// BFS to find shortest path from 0 to n-1
  // Alokasi slice integer
		dist := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
		for i := range dist {
			dist[i] = -1
		}
		dist[0] = 0
		q := list.New()
		q.PushBack(0)

		for q.Len() > 0 {
			u := q.Remove(q.Front()).(int)
			if u == n-1 {
				break
			}
			for _, v := range adj[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					q.PushBack(v)
				}
			}
		}

		if dist[n-1] != -1 && dist[n-1] <= k {
			ans = costLimit
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(MinimumCostToRepairEdgesToTraverseAGraph(3, [][]int{{0, 1, 10}, {1, 2, 10}, {0, 2, 100}}, 1)) // Expected: 100

	// Example 2
	edges2 := [][]int{{0, 2, 5}, {2, 3, 6}, {3, 4, 7}, {4, 5, 5}, {0, 1, 10}, {1, 5, 12}, {0, 3, 9}, {1, 2, 8}, {2, 4, 11}}
	fmt.Println(MinimumCostToRepairEdgesToTraverseAGraph(6, edges2, 2)) // Expected: 12

	// Example 3
	fmt.Println(MinimumCostToRepairEdgesToTraverseAGraph(3, [][]int{{0, 1, 1}}, 1)) // Expected: -1
}
```
