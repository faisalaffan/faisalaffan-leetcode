# 1615 — Maximal Network Rank

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func MaximalNetworkRank(n int, roads [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(N^2), Space: O(N^2)  |  **Ruang:** O(N^2)


## 💻 Solusi Go

```go
package main

// LeetCode #1615: Maximal Network Rank
// https://leetcode.com/problems/maximal-network-rank/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaximalNetworkRank(4, [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}}))
	fmt.Println(MaximalNetworkRank(5, [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}, {2, 3}, {2, 4}}))
	fmt.Println(MaximalNetworkRank(2, [][]int{{0, 1}}))
}

func MaximalNetworkRank(n int, roads [][]int) int {
	// Time: O(N^2), Space: O(N^2)
  // Alokasi slice
	degree := make([]int, n)
  // Matriks 2D
	connected := make([][]bool, n)
	for i := 0; i < n; i++ {
		connected[i] = make([]bool, n)
	}

	for _, r := range roads {
		u, v := r[0], r[1]
		degree[u]++
		degree[v]++
		connected[u][v] = true
		connected[v][u] = true
	}

	maxRank := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			rank := degree[i] + degree[j]
			if connected[i][j] {
				rank-- // shared edge counted twice
			}
			if rank > maxRank {
				maxRank = rank
			}
		}
	}

	return maxRank
}
```
