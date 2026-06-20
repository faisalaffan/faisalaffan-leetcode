# 2556 — Disconnect Path In A Binary Matrix By At Most One Flip

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func isPossibleToCutPath(grid [][]int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(m*n)  |  **Ruang:** O(m*n)


## 💻 Solusi Go

```go
package main

// LeetCode #2556: Disconnect Path in a Binary Matrix by at Most One Flip
// https://leetcode.com/problems/disconnect-path-in-a-binary-matrix-by-at-most-one-flip/
// Difficulty: Medium
// Time: O(m*n) | Space: O(m*n)

import "fmt"

func isPossibleToCutPath(grid [][]int) bool {
	m, n := len(grid), len(grid[0])

	// First DFS from (0,0) to (m-1,n-1), mark visited cells
	var dfs1 func(r, c int) bool
	dfs1 = func(r, c int) bool {
		if r >= m || c >= n || grid[r][c] == 0 {
			return false
		}
		if r == m-1 && c == n-1 {
			return true
		}
		grid[r][c] = 0 // Mark as visited
		return dfs1(r+1, c) || dfs1(r, c+1)
	}

	if !dfs1(0, 0) {
		return true // Already disconnected
	}

	// Second DFS from (0,0) checking if there's still a path after first removal
	var dfs2 func(r, c int) bool
	dfs2 = func(r, c int) bool {
		if r >= m || c >= n || grid[r][c] == 0 {
			return false
		}
		if r == m-1 && c == n-1 {
			return true
		}
		grid[r][c] = 0
		return dfs2(r, c+1) || dfs2(r+1, c)
	}

	if !dfs2(0, 0) {
		return true
	}
	return false
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", isPossibleToCutPath([][]int{{1, 1, 1}, {1, 0, 0}, {1, 1, 1}}))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", isPossibleToCutPath([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}))
	// Expected: false

	// Test case 3: single cell
	fmt.Println("Test 3:", isPossibleToCutPath([][]int{{1}}))
	// Expected: false
}
```
