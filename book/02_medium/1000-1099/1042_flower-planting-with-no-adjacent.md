# 1042 — Flower Planting With No Adjacent

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func gardenNoAdj(n int, paths [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n + e) where e = len(paths)  |  **Ruang:** O(n + e)


## 💻 Solusi Go

```go
package main

// LeetCode #1042: Flower Planting With No Adjacent
// https://leetcode.com/problems/flower-planting-with-no-adjacent/
// Difficulty: Medium
//
// Approach: Greedy coloring with a graph. Each garden gets a flower type
//           not used by its neighbors.
// Time: O(n + e) where e = len(paths)
// Space: O(n + e)

import "fmt"

func main() {
	fmt.Println(gardenNoAdj(3, [][]int{{1, 2}, {2, 3}, {3, 1}})) // [1,2,3]
	fmt.Println(gardenNoAdj(4, [][]int{{1, 2}, {3, 4}}))         // [1,2,1,2]
}

func gardenNoAdj(n int, paths [][]int) []int {
  // Matriks 2D
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}

	for _, p := range paths {
		u, v := p[0]-1, p[1]-1
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

  // Alokasi slice
	result := make([]int, n)
	for i := 0; i < n; i++ {
		used := make([]bool, 5)
		for _, neighbor := range graph[i] {
			if result[neighbor] != 0 {
				used[result[neighbor]] = true
			}
		}
		for flower := 1; flower <= 4; flower++ {
			if !used[flower] {
				result[i] = flower
				break
			}
		}
	}

	return result
}
```
