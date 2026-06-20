# 2101 — Detonate The Maximum Bombs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func maximumDetonation(bombs [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** BFS

**Waktu:** O(n^2)  |  **Ruang:** O(n^2)

> 🎓 **Fresh Grad Tips:** Kuasai **BFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2101: Detonate the Maximum Bombs
// https://leetcode.com/problems/detonate-the-maximum-bombs/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n^2)

import "fmt"

func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
  // Edge case: input kosong
	if n == 0 {
		return 0
	}

	// Build adjacency (directed from i to j if i can detonate j)
  // Matriks 2D
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				dx := bombs[i][0] - bombs[j][0]
				dy := bombs[i][1] - bombs[j][1]
				dist := int64(dx)*int64(dx) + int64(dy)*int64(dy)
				radius := int64(bombs[i][2])
				if dist <= radius*radius {
					adj[i] = append(adj[i], j)
				}
			}
		}
	}

	maxDetonated := 0
	for i := 0; i < n; i++ {
		visited := make([]bool, n)
		queue := []int{i}
		visited[i] = true
		count := 0

		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			count++
			for _, next := range adj[curr] {
				if !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		}
		if count > maxDetonated {
			maxDetonated = count
		}
	}
	return maxDetonated
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maximumDetonation([][]int{{2, 1, 3}, {6, 1, 4}}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", maximumDetonation([][]int{{1, 1, 5}, {10, 10, 5}}))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", maximumDetonation([][]int{{1, 2, 3}, {2, 3, 1}, {3, 4, 2}, {4, 5, 3}, {5, 6, 4}}))
	// Expected: 5
}
```
