# 1931 — Painting A Grid With Three Different Colors

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func colorTheGrid(m int, n int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1931: Painting a Grid With Three Different Colors
// https://leetcode.com/problems/painting-a-grid-with-three-different-colors/
// Difficulty: Hard
// DP on columns, m <= 5, 3 colors. Encode each column as base-3 integer.
// dp[col][mask] = ways.

import "fmt"

const mod = 1_000_000_007

func colorTheGrid(m int, n int) int {
	// Generate all valid column patterns (no adjacent cells same color in a column)
	total := 1
	for i := 0; i < m; i++ {
		total *= 3
	}

	var masks []int
	for mask := 0; mask < total; mask++ {
		valid := true
		prev := -1
		tmp := mask
		for i := 0; i < m; i++ {
			cur := tmp % 3
			tmp /= 3
			if cur == prev {
				valid = false
				break
			}
			prev = cur
		}
		if valid {
			masks = append(masks, mask)
		}
	}

	// Precompute transitions: check if two columns can be adjacent
	canTransition := func(a, b int) bool {
		for i := 0; i < m; i++ {
			if a%3 == b%3 {
				return false
			}
			a /= 3
			b /= 3
		}
		return true
	}

	// dp for current column
  // Alokasi slice
	dp := make([]int, total)
	for _, mask := range masks {
		dp[mask] = 1
	}

	for col := 1; col < n; col++ {
  // Alokasi slice
		ndp := make([]int, total)
		for _, a := range masks {
			if dp[a] == 0 {
				continue
			}
			for _, b := range masks {
				if canTransition(a, b) {
					ndp[b] = (ndp[b] + dp[a]) % mod
				}
			}
		}
		dp = ndp
	}

	ans := 0
	for _, mask := range masks {
		ans = (ans + dp[mask]) % mod
	}
	return ans
}

func main() {
	fmt.Println(colorTheGrid(1, 1)) // Expected: 3
	fmt.Println(colorTheGrid(1, 2)) // Expected: 6
	fmt.Println(colorTheGrid(2, 1)) // Expected: 6
	fmt.Println(colorTheGrid(2, 2)) // Expected: 18
	fmt.Println(colorTheGrid(5, 1000)) // Expected: something (performance test)
}
```
