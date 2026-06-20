# 3820 — Pythagorean Distance Nodes In A Tree

## Deskripsi

**Soal:** [3820. Pythagorean Distance Nodes In A Tree](https://leetcode.com/problems/pythagorean-distance-nodes-in-a-tree/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(N)

**Algoritma:** BFS (Breadth-First Search / pencarian lebar), LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func PythagoreanDistanceNodesInATree(n int, edges [][]int, x int, y int, z int) int`

> **Ide Kunci:** BFS from each target node x, y, z to compute distances,

## Solusi Go

```go
package main

// LeetCode #3820: Pythagorean Distance Nodes in a Tree
// https://leetcode.com/problems/pythagorean-distance-nodes-in-a-tree/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: BFS from each target node x, y, z to compute distances,
// then count nodes where sorted distances form a Pythagorean triple.

import (
	"fmt"
	"sort"
)

func PythagoreanDistanceNodesInATree(n int, edges [][]int, x int, y int, z int) int {
	// Build adjacency list
  // Membuat slice 2D untuk DP/tabel
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// BFS to compute distances from a source
	bfs := func(src int) []int {
  // Membuat slice untuk menyimpan hasil
		dist := make([]int, n)
  // Iterasi seluruh elemen
		for i := range dist {
			dist[i] = -1
		}
		dist[src] = 0
		q := []int{src}
		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			for _, v := range adj[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					q = append(q, v)
				}
			}
		}
		return dist
	}

	distX := bfs(x)
	distY := bfs(y)
	distZ := bfs(z)

	ans := 0
	for i := 0; i < n; i++ {
		d := []int{distX[i], distY[i], distZ[i]}
		sort.Ints(d)
		a, b, c := d[0], d[1], d[2]
		if a*a+b*b == c*c {
			ans++
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(PythagoreanDistanceNodesInATree(4, [][]int{{0, 1}, {0, 2}, {0, 3}}, 1, 2, 3)) // Expected: 3

	// Example 2
	fmt.Println(PythagoreanDistanceNodesInATree(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, 0, 2, 4)) // Expected: ?

	// Example 3
	fmt.Println(PythagoreanDistanceNodesInATree(3, [][]int{{0, 1}, {1, 2}}, 0, 1, 2))
}
```
