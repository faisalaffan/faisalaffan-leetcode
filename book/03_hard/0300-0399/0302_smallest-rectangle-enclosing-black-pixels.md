# 0302 — Smallest Rectangle Enclosing Black Pixels

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minArea(image [][]byte, x int, y int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Binary Search

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #302: Smallest Rectangle Enclosing Black Pixels
// https://leetcode.com/problems/smallest-rectangle-enclosing-black-pixels/
// Difficulty: Hard [Paid]
//
// Approach: Binary Search on Boundaries.
//   The image has exactly one black region (all '1's are contiguous).
//   Given a known black pixel (x, y), we can binary search for the
//   topmost, bottommost, leftmost, and rightmost '1's.
//   - Top: smallest row in [0, x] that has a '1' in any column.
//   - Bottom: largest row in [x, m-1] that has a '1' in any column.
//   - Left: smallest col in [0, y] that has a '1' in any row.
//   - Right: largest col in [y, n-1] that has a '1' in any row.
//   Area = (bottom - top + 1) * (right - left + 1).

import "fmt"

func main() {
	// Example:
	// [["0","0","1","0"],
	//  ["0","1","1","0"],
	//  ["0","1","0","0"]]
	// x=0, y=2 -> area=6
	grid := [][]byte{
		{'0', '0', '1', '0'},
		{'0', '1', '1', '0'},
		{'0', '1', '0', '0'},
	}
	fmt.Println("Area:", minArea(grid, 0, 2)) // 6

	// Single pixel.
	grid2 := [][]byte{{'1'}}
	fmt.Println("Area (single):", minArea(grid2, 0, 0)) // 1

	// Single column.
	grid3 := [][]byte{
		{'0', '1'},
		{'0', '1'},
	}
	fmt.Println("Area (col):", minArea(grid3, 0, 1)) // 2

	// All zeros (edge case, shouldn't happen per problem constraints).
	// The problem guarantees at least one black pixel at (x,y).
	grid4 := [][]byte{
		{'0', '0'},
		{'1', '0'},
	}
	fmt.Println("Area (one pixel):", minArea(grid4, 1, 0)) // 1
}

// minArea returns the area of the smallest rectangle that encloses all '1's.
func minArea(image [][]byte, x int, y int) int {
	if len(image) == 0 || len(image[0]) == 0 {
		return 0
	}

	m, n := len(image), len(image[0])

	// Binary search for topmost row with a '1'.
	top := searchTop(image, 0, x, n)
	// Binary search for bottommost row with a '1'.
	bottom := searchBottom(image, x, m-1, n)
	// Binary search for leftmost column with a '1'.
	left := searchLeft(image, 0, y, m)
	// Binary search for rightmost column with a '1'.
	right := searchRight(image, y, n-1, m)

	return (bottom - top + 1) * (right - left + 1)
}

// hasBlackInRow checks if any column in row r has a '1'.
func hasBlackInRow(image [][]byte, row int, n int) bool {
	for c := 0; c < n; c++ {
		if image[row][c] == '1' {
			return true
		}
	}
	return false
}

// hasBlackInColumn checks if any row in column c has a '1'.
func hasBlackInColumn(image [][]byte, col int, m int) bool {
	for r := 0; r < m; r++ {
		if image[r][col] == '1' {
			return true
		}
	}
	return false
}

// searchTop finds the topmost row with a '1' in range [lo, hi].
func searchTop(image [][]byte, lo, hi int, n int) int {
	for lo < hi {
		mid := lo + (hi-lo)/2
		if hasBlackInRow(image, mid, n) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// searchBottom finds the bottommost row with a '1' in range [lo, hi].
func searchBottom(image [][]byte, lo, hi int, n int) int {
	for lo < hi {
		mid := lo + (hi-lo+1)/2 // ceiling mid
		if hasBlackInRow(image, mid, n) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}

// searchLeft finds the leftmost column with a '1' in range [lo, hi].
func searchLeft(image [][]byte, lo, hi int, m int) int {
	for lo < hi {
		mid := lo + (hi-lo)/2
		if hasBlackInColumn(image, mid, m) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// searchRight finds the rightmost column with a '1' in range [lo, hi].
func searchRight(image [][]byte, lo, hi int, m int) int {
	for lo < hi {
		mid := lo + (hi-lo+1)/2 // ceiling mid
		if hasBlackInColumn(image, mid, m) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}

// Stub compatibility.
func SmallestRectangleEnclosingBlackPixels() any {
	grid := [][]byte{
		{'0', '0', '1', '0'},
		{'0', '1', '1', '0'},
		{'0', '1', '0', '0'},
	}
	return minArea(grid, 0, 2)
}
```
