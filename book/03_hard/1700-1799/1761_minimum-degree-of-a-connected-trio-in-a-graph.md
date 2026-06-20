# 1761 — Minimum Degree Of A Connected Trio In A Graph

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func minTrioDegree(n int, edges [][]int) int
```

> **💡 Hint:** Adjacency matrix + degree array.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1761: Minimum Degree of a Connected Trio in a Graph
// https://leetcode.com/problems/minimum-degree-of-a-connected-trio-in-a-graph/
// Difficulty: Hard
//
// Approach: Adjacency matrix + degree array.
// 1. Build degree array and adjacency matrix for the graph.
// 2. For every trio (i, j, k) with i < j < k, check if all three edges exist.
// 3. Degree of trio = degree[i] + degree[j] + degree[k] - 6 (each internal edge counted twice).
// 4. Track minimum.

import (
	"fmt"
	"math"
)

func minTrioDegree(n int, edges [][]int) int {
  // Alokasi slice integer
	deg := make([]int, n)
  // Membuat matriks/slice 2D untuk DP
	adj := make([][]bool, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range adj {
		adj[i] = make([]bool, n)
	}

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		deg[u]++
		deg[v]++
		adj[u][v] = true
		adj[v][u] = true
	}

	minDeg := math.MaxInt32
	found := false

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if !adj[i][j] {
				continue
			}
			for k := j + 1; k < n; k++ {
				if adj[i][k] && adj[j][k] {
					found = true
					d := deg[i] + deg[j] + deg[k] - 6
					if d < minDeg {
						minDeg = d
					}
				}
			}
		}
	}

	if !found {
		return -1
	}
	return minDeg
}

func main() {
	// Example test case
	n := 6
	edges := [][]int{{1, 2}, {1, 3}, {3, 2}, {4, 1}, {5, 2}, {3, 6}}
	fmt.Println("n=6,edges=[[1,2],[1,3],[3,2],[4,1],[5,2],[3,6]] →", minTrioDegree(n, edges)) // Expected: 3

	// Additional tests
	n2 := 4
	edges2 := [][]int{{1, 2}, {1, 3}, {3, 2}}
	fmt.Println("n=4,edges=[[1,2],[1,3],[3,2]] →", minTrioDegree(n2, edges2)) // Expected: 2 (deg=2,2,2 => 2+2+2-6=0? No deg[1]=2, deg[2]=2, deg[3]=2 => 0)

	n3 := 3
	edges3 := [][]int{{1, 2}, {2, 3}, {1, 3}}
	fmt.Println("n=3,edges=[[1,2],[2,3],[1,3]] →", minTrioDegree(n3, edges3)) // Expected: 0

	n4 := 5
	edges4 := [][]int{{1, 2}, {2, 3}, {3, 1}, {1, 4}, {2, 5}}
	fmt.Println("n=5,edges=[[1,2],[2,3],[3,1],[1,4],[2,5]] →", minTrioDegree(n4, edges4)) // Expected: 2
}
```
