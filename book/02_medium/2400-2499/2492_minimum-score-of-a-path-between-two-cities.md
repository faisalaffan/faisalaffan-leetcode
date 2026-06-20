# 2492 — Minimum Score Of A Path Between Two Cities

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func minScore(n int, roads [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n + m)  
**Kompleksitas Ruang:** O(n + m)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2492: Minimum Score of a Path Between Two Cities
// https://leetcode.com/problems/minimum-score-of-a-path-between-two-cities/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n + m)
// DFS from city 1, find min edge in its connected component (must include city n).

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minScore(4, [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}})) // 5
	fmt.Println(minScore(4, [][]int{{1, 2, 2}, {1, 3, 4}, {3, 4, 7}}))            // 2
}

func minScore(n int, roads [][]int) int {
  // Membuat matriks/slice 2D untuk DP
	graph := make([][][2]int, n+1)
	for _, r := range roads {
		a, b, d := r[0], r[1], r[2]
		graph[a] = append(graph[a], [2]int{b, d})
		graph[b] = append(graph[b], [2]int{a, d})
	}

	visited := make([]bool, n+1)
	ans := math.MaxInt32

	var dfs func(u int)
	dfs = func(u int) {
		visited[u] = true
		for _, edge := range graph[u] {
			v, d := edge[0], edge[1]
			if d < ans {
				ans = d
			}
			if !visited[v] {
				dfs(v)
			}
		}
	}
	dfs(1)

	return ans
}
```
