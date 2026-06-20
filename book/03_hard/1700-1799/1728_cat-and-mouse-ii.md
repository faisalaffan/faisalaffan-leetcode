# 1728 — Cat And Mouse Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func canMouseWin(grid []string, catJump int, mouseJump int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1728: Cat and Mouse II
// https://leetcode.com/problems/cat-and-mouse-ii/
// Difficulty: Hard

import "fmt"

func main() {
	grid1 := []string{
		"####F",
		"#C...#",
		"M....#",
	}
	fmt.Printf("Test 1:\nResult: %v (Expected: true)\n\n", canMouseWin(grid1, 1, 2))

	grid2 := []string{
		"M.C...F",
	}
	fmt.Printf("Test 2:\nResult: %v (Expected: true)\n\n", canMouseWin(grid2, 1, 4))

	grid3 := []string{
		"M.C...F",
	}
	fmt.Printf("Test 3:\nResult: %v (Expected: false)\n\n", canMouseWin(grid3, 1, 3))

	grid4 := []string{
		"C...#",
		"...#F",
		"....#",
		"M....",
	}
	fmt.Printf("Test 4:\nResult: %v (Expected: false)\n", canMouseWin(grid4, 2, 5))
}

func canMouseWin(grid []string, catJump int, mouseJump int) bool {
	rows := len(grid)
	cols := len(grid[0])

	// Locate positions
	var mr, mc, cr, cc, fr, fc int
	found := 0
	for i := 0; i < rows && found < 6; i++ {
		for j := 0; j < cols && found < 6; j++ {
			switch grid[i][j] {
			case 'M':
				mr, mc = i, j
				found++
			case 'C':
				cr, cc = i, j
				found++
			case 'F':
				fr, fc = i, j
				found++
			}
		}
	}

	maxMoves := rows * cols * 2
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

  // HashMap: O(1) lookup
	memo := make(map[[5]int]bool)

	var dfs func(mr, mc, cr, cc int, turn int, moves int) bool
	dfs = func(mr, mc, cr, cc int, turn int, moves int) bool {
		if moves > maxMoves {
			return false
		}
		key := [5]int{mr, mc, cr, cc, turn}
		if res, ok := memo[key]; ok {
			return res
		}

		if turn == 0 { // Mouse's turn
			// Option: stay in place
			if mr == fr && mc == fc {
				memo[key] = true
				return true
			}
			if dfs(mr, mc, cr, cc, 1, moves+1) {
				memo[key] = true
				return true
			}

			// Option: move in each direction
			for _, d := range dirs {
				for step := 1; step <= mouseJump; step++ {
					nr, nc := mr+d[0]*step, mc+d[1]*step
					if nr < 0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] == '#' {
						break
					}
					if nr == cr && nc == cc {
						continue
					}
					if nr == fr && nc == fc {
						memo[key] = true
						return true
					}
					if dfs(nr, nc, cr, cc, 1, moves+1) {
						memo[key] = true
						return true
					}
				}
			}
			memo[key] = false
			return false
		} else { // Cat's turn
			// Option: stay in place
			if cr == mr && cc == mc {
				memo[key] = false
				return false
			}
			if cr == fr && cc == fc {
				memo[key] = false
				return false
			}
			if !dfs(mr, mc, cr, cc, 0, moves+1) {
				memo[key] = false
				return false
			}

			// Option: move in each direction
			for _, d := range dirs {
				for step := 1; step <= catJump; step++ {
					nr, nc := cr+d[0]*step, cc+d[1]*step
					if nr < 0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] == '#' {
						break
					}
					if nr == mr && nc == mc {
						memo[key] = false
						return false
					}
					if nr == fr && nc == fc {
						memo[key] = false
						return false
					}
					if !dfs(mr, mc, nr, nc, 0, moves+1) {
						memo[key] = false
						return false
					}
				}
			}
			memo[key] = true
			return true
		}
	}

	return dfs(mr, mc, cr, cc, 0, 0)
}
```
