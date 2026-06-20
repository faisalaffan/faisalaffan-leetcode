# 3609 — Minimum Moves To Reach Target In Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minMoves(sx, sy, tx, ty int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #3609: Minimum Moves to Reach Target in Grid
// https://leetcode.com/problems/minimum-moves-to-reach-target-in-grid/
// Difficulty: Hard
//
// On an infinite grid, from (x, y) you can go to (x+m, y) or (x, y+m) where
// m = max(x, y). Find minimum moves from (sx, sy) to (tx, ty).
//
// Approach: Work backwards from (tx, ty) to (sx, sy) using reverse operations.

import "fmt"

func main() {
	// Example 1
	fmt.Println(minMoves(1, 1, 4, 3))
	// Example 2
	fmt.Println(minMoves(0, 0, 3, 4))
	// Edge: already at target
	fmt.Println(minMoves(5, 3, 5, 3))
	// Edge: impossible
	fmt.Println(minMoves(1, 2, 3, 3))
}

func minMoves(sx, sy, tx, ty int) int {
	moves := 0
	for sx < tx || sy < ty {
		if tx == ty {
			return -1
		}
		if tx > ty {
			// Forward move added to x
			// If tx >= 2*ty: came from (tx/2, ty) via doubling
			// Else: came from (tx-ty, ty) via addition
			if ty < sy {
				return -1
			}
			if sy == ty {
				// y is fixed, only x changes
				if (tx-sx)%ty == 0 {
					return moves + (tx-sx)/ty
				}
				return -1
			}
			if tx >= 2*ty {
				tx /= 2
			} else {
				tx -= ty
			}
		} else {
			if tx < sx {
				return -1
			}
			if sx == tx {
				if (ty-sy)%tx == 0 {
					return moves + (ty-sy)/tx
				}
				return -1
			}
			if ty >= 2*tx {
				ty /= 2
			} else {
				ty -= tx
			}
		}
		moves++
	}
	return moves
}
```
