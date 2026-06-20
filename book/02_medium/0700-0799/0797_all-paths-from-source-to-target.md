# 0797 — All Paths From Source To Target

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func allPathsSourceTarget(graph [][]int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DFS

**Waktu:** O(2^n * n)  |  **Ruang:** O(2^n * n)

> 🎓 **Fresh Grad Tips:** Kuasai **DFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #797: All Paths From Source to Target
// https://leetcode.com/problems/all-paths-from-source-to-target/
// Difficulty: Medium
// Time: O(2^n * n)
// Space: O(2^n * n)

import "fmt"

func main() {
	fmt.Println(allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))
	fmt.Println(allPathsSourceTarget([][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}))
}

func allPathsSourceTarget(graph [][]int) [][]int {
  // Matriks 2D
	result := make([][]int, 0)
  // Alokasi slice
	path := make([]int, 0)
	path = append(path, 0)

	var dfs func(node int)
	dfs = func(node int) {
		if node == len(graph)-1 {
  // Alokasi slice
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			result = append(result, pathCopy)
			return
		}

		for _, neighbor := range graph[node] {
			path = append(path, neighbor)
			dfs(neighbor)
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	return result
}
```
