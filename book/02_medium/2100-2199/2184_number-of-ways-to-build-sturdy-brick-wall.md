# 2184 — Number Of Ways To Build Sturdy Brick Wall

## Deskripsi

**Soal:** [2184. Number Of Ways To Build Sturdy Brick Wall](https://leetcode.com/problems/number-of-ways-to-build-sturdy-brick-wall/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * m)  
**Kompleksitas Ruang:** O(w) where w = max width

**Algoritma:** Dynamic Programming (DP), DFS (Depth-First Search / pencarian kedalaman)

**Fungsi Solusi:** `func buildWall(height int, width int, bricks []int) int`

## Solusi Go

```go
package main

// LeetCode #2184: Number of Ways to Build Sturdy Brick Wall
// https://leetcode.com/problems/number-of-ways-to-build-sturdy-brick-wall/
// Difficulty: Medium [Paid]
// Time: O(n * m) | Space: O(w) where w = max width

import "fmt"

func buildWall(height int, width int, bricks []int) int {
	const mod = 1_000_000_007

	// Generate all valid row patterns via DFS
	var rows [][]int
	var dfs func(curr []int, w int)
	dfs = func(curr []int, w int) {
		if w == width {
  // Membuat slice untuk menyimpan hasil
			row := make([]int, len(curr))
			copy(row, curr)
			rows = append(rows, row)
			return
		}
		for _, b := range bricks {
			if w+b <= width {
				dfs(append(curr, b), w+b)
			}
		}
	}
	dfs([]int{}, 0)

	// Precompute which rows are compatible (no shared seam)
	n := len(rows)
  // Membuat slice 2D untuk DP/tabel
	compat := make([][]bool, n)
	for i := 0; i < n; i++ {
		compat[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			compat[i][j] = true
		}
	}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			seami, seamj := 0, 0
			pi, pj := 0, 0
			for pi < len(rows[i])-1 && pj < len(rows[j])-1 {
				seami += rows[i][pi]
				seamj += rows[j][pj]
				if seami == seamj {
					compat[i][j] = false
					compat[j][i] = false
					break
				}
				if seami < seamj {
					pi++
				} else {
					pj++
				}
			}
		}
	}

	// DP: ways[h][r] = ways to build up to height h ending with row r
  // Membuat slice 2D untuk DP/tabel
	dp := make([][]int, height)
	for h := 0; h < height; h++ {
		dp[h] = make([]int, n)
	}
	for r := 0; r < n; r++ {
		dp[0][r] = 1
	}

	for h := 1; h < height; h++ {
		for r := 0; r < n; r++ {
			for p := 0; p < n; p++ {
				if compat[p][r] {
					dp[h][r] = (dp[h][r] + dp[h-1][p]) % mod
				}
			}
		}
	}

	total := 0
	for r := 0; r < n; r++ {
		total = (total + dp[height-1][r]) % mod
	}
	return total
}

func main() {
	// Test case 1
	fmt.Println(buildWall(2, 3, []int{1, 2}))
	// Expected: 2

	// Test case 2
	fmt.Println(buildWall(1, 1, []int{1}))
	// Expected: 1
}
```
