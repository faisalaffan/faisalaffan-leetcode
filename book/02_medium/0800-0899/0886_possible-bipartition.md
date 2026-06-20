# 0886 — Possible Bipartition

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func PossibleBipartition(n int, dislikes [][]int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n + d) where d = len(dislikes)  |  **Ruang:** O(n + d)


## 💻 Solusi Go

```go
package main

// LeetCode #886: Possible Bipartition
// https://leetcode.com/problems/possible-bipartition/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(PossibleBipartition(4, [][]int{{1, 2}, {1, 3}, {2, 4}}))
	fmt.Println(PossibleBipartition(3, [][]int{{1, 2}, {1, 3}, {2, 3}}))
	fmt.Println(PossibleBipartition(5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}}))
}

// Time: O(n + d) where d = len(dislikes) | Space: O(n + d)
func PossibleBipartition(n int, dislikes [][]int) bool {
  // Matriks 2D
	graph := make([][]int, n+1)
	for _, d := range dislikes {
		a, b := d[0], d[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

  // Alokasi slice
	color := make([]int, n+1) // 0 = uncolored, 1 = group A, -1 = group B

	var dfs func(node, c int) bool
	dfs = func(node, c int) bool {
		if color[node] != 0 {
			return color[node] == c
		}
		color[node] = c
		for _, nei := range graph[node] {
			if !dfs(nei, -c) {
				return false
			}
		}
		return true
	}

	for i := 1; i <= n; i++ {
		if color[i] == 0 && !dfs(i, 1) {
			return false
		}
	}

	return true
}
```
