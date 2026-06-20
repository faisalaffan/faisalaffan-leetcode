# 0711 — Number Of Distinct Islands Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func numDistinctIslandsII(grid [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Two Pointer, DFS, Sorting

**Waktu:** O(rows*cols * log(rows*cols)) for island detection and normalization  |  **Ruang:** O(rows*cols)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #711: Number of Distinct Islands II
// https://leetcode.com/problems/number-of-distinct-islands-ii/
// Difficulty: Hard [Paid]
//
// Given a 2D grid of 1s (land) and 0s (water), count the number of distinct islands.
// Two islands are considered the same if one can be translated, rotated (90°, 180°, 270°),
// or reflected (mirrored) to match the other. All 8 transformations are considered.

// pair represents a cell coordinate (row, col).
type pair struct {
	r, c int
}

// numDistinctIslandsII counts distinct islands considering all 8 transformations.
// Time: O(rows*cols * log(rows*cols)) for island detection and normalization
// Space: O(rows*cols)
func numDistinctIslandsII(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
  // Matriks 2D
	visited := make([][]bool, rows)
  // Range loop
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	// All 8 transformations: (newR, newC) = f(r, c)
	// Each transform is represented as (a, b, c, d) where:
	//   newR = a*r + b*c
	//   newC = c*r + d*c
	// Encoding: index -> (a,b,c,d)
	// 0: (1,0,0,1)  identity
	// 1: (1,0,0,-1) reflect over x-axis
	// 2: (-1,0,0,1) reflect over y-axis
	// 3: (-1,0,0,-1) 180° rotation
	// 4: (0,1,1,0) reflect over y=x (swap)
	// 5: (0,1,-1,0) 90° CCW
	// 6: (0,-1,1,0) 270° CCW
	// 7: (0,-1,-1,0) reflect over y=-x
	transforms := [][4]int{
		{1, 0, 0, 1},   // 0: identity
		{1, 0, 0, -1},  // 1: reflect x
		{-1, 0, 0, 1},  // 2: reflect y
		{-1, 0, 0, -1}, // 3: rotate 180
		{0, 1, 1, 0},   // 4: reflect y=x
		{0, 1, -1, 0},  // 5: rotate 90 CCW
		{0, -1, 1, 0},  // 6: rotate 270 CCW
		{0, -1, -1, 0}, // 7: reflect y=-x
	}

	// canonicalForm generates the canonical (normalized) representation of a shape
	// under all 8 transformations. It returns a string key.
	canonicalForm := func(shape []pair) string {
		if len(shape) == 0 {
			return ""
		}

		// Try all 8 transformations and pick the lexicographically smallest.
		var bestStr string

		for _, t := range transforms {
			// Apply transformation.
			transformed := make([]pair, len(shape))
			for i, p := range shape {
				newR := t[0]*p.r + t[1]*p.c
				newC := t[2]*p.r + t[3]*p.c
				transformed[i] = pair{newR, newC}
			}

			// Sort transformed points.
  // Custom sort
			sort.Slice(transformed, func(i, j int) bool {
				if transformed[i].r != transformed[j].r {
					return transformed[i].r < transformed[j].r
				}
				return transformed[i].c < transformed[j].c
			})

			// Translate to origin (subtract min r, min c).
			minR, minC := transformed[0].r, transformed[0].c
  // Range loop
			for i := range transformed {
				transformed[i].r -= minR
				transformed[i].c -= minC
			}

			// Build string representation.
			var key string
			for _, p := range transformed {
				key += fmt.Sprintf("(%d,%d)", p.r, p.c)
			}

			if bestStr == "" || key < bestStr {
				bestStr = key
			}
		}
		return bestStr
	}

	// DFS to find all cells of an island.
	var dfs func(r, c int, shape *[]pair)
	dfs = func(r, c int, shape *[]pair) {
		if r < 0 || r >= rows || c < 0 || c >= cols || visited[r][c] || grid[r][c] == 0 {
			return
		}
		visited[r][c] = true
		*shape = append(*shape, pair{r, c})
		dfs(r-1, c, shape)
		dfs(r+1, c, shape)
		dfs(r, c-1, shape)
		dfs(r, c+1, shape)
	}

  // HashMap: O(1) lookup
	islandSet := make(map[string]bool)

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 && !visited[r][c] {
				var shape []pair
				dfs(r, c, &shape)
				key := canonicalForm(shape)
				islandSet[key] = true
			}
		}
	}

	return len(islandSet)
}

// --- Tests -------------------------------------------------------------------

func main() {
	fmt.Println("=== 0711 Number of Distinct Islands II ===")

	// Test 1: Two identical islands (same rotation).
	// Island 1: (0,0)-(0,1)-(1,0)
	// Island 2: (2,2)-(2,3)-(3,2) - same shape, just translated
	grid1 := [][]int{
		{1, 1, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 1, 0, 0},
	}
	n1 := numDistinctIslandsII(grid1)
	fmt.Printf("Test 1 - Distinct islands = %d (expected 1)\n", n1)

	// Test 2: Two distinct islands (different shapes).
	// Island 1: L-shape (0,0)-(0,1)-(1,0)
	// Island 2: straight line (0,3)-(1,3)-(2,3)
	grid2 := [][]int{
		{1, 1, 0, 1},
		{1, 0, 0, 1},
		{0, 0, 0, 1},
	}
	n2 := numDistinctIslandsII(grid2)
	fmt.Printf("Test 2 - Distinct islands = %d (expected 2)\n", n2)

	// Test 3: Islands that are rotations of each other count as the same.
	// Left L: (0,0)-(0,1)-(1,0)
	// Right rotated L: (0,3)-(1,3)-(1,4) which is the L rotated 90 deg
	grid3 := [][]int{
		{1, 1, 0, 0, 1},
		{1, 0, 0, 1, 1},
	}
	n3 := numDistinctIslandsII(grid3)
	fmt.Printf("Test 3 - Distinct islands = %d (expected 1 - rotations count as same)\n", n3)

	// Test 4: Single cell islands are all the same.
	grid4 := [][]int{
		{1, 0, 1},
		{0, 0, 0},
		{1, 0, 1},
	}
	n4 := numDistinctIslandsII(grid4)
	fmt.Printf("Test 4 - Single cells = %d (expected 1)\n", n4)

	// Test 5: Empty grid.
	n5 := numDistinctIslandsII([][]int{})
	fmt.Printf("Test 5 - Empty grid = %d (expected 0)\n", n5)
}
```
