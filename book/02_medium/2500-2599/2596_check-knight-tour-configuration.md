# 2596 — Check Knight Tour Configuration

## Deskripsi

**Soal:** [2596. Check Knight Tour Configuration](https://leetcode.com/problems/check-knight-tour-configuration/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n^2)

**Algoritma:** —

**Fungsi Solusi:** `func checkValidGrid(grid [][]int) bool`

## Solusi Go

```go
package main

// LeetCode #2596: Check Knight Tour Configuration
// https://leetcode.com/problems/check-knight-tour-configuration/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n^2)

import "fmt"

func checkValidGrid(grid [][]int) bool {
	n := len(grid)
  // Edge case: input kosong
	if n == 0 {
		return false
	}
	if grid[0][0] != 0 {
		return false
	}

	// Find positions of each move
  // Membuat slice untuk menyimpan hasil
	pos := make([][2]int, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			pos[grid[i][j]] = [2]int{i, j}
		}
	}

	// Knight moves
	dirs := [8][2]int{{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {2, -1}, {2, 1}}

	for k := 1; k < n*n; k++ {
		r, c := pos[k][0], pos[k][1]
		pr, pc := pos[k-1][0], pos[k-1][1]
		valid := false
		for _, d := range dirs {
			if pr+d[0] == r && pc+d[1] == c {
				valid = true
				break
			}
		}
		if !valid {
			return false
		}
	}
	return true
}

func main() {
	// Test case 1: valid
	fmt.Println("Test 1:", checkValidGrid([][]int{{0, 11, 16, 5, 20}, {17, 4, 19, 10, 15}, {12, 1, 8, 21, 6}, {3, 18, 23, 14, 9}, {24, 13, 2, 7, 22}}))
	// Expected: true

	// Test case 2: invalid (grid[0][0] != 0)
	fmt.Println("Test 2:", checkValidGrid([][]int{{1, 0}, {0, 1}}))
	// Expected: false

	// Test case 3: 1x1 grid
	fmt.Println("Test 3:", checkValidGrid([][]int{{0}}))
	// Expected: true
}
```
