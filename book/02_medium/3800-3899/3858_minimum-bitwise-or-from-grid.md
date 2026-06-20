# 3858 — Minimum Bitwise Or From Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumBitwiseOrFromGrid(grid [][]int) int
```

> **💡 Hint:** Greedy bit-by-bit from MSB. Try to keep each bit as 0 if

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(32 * M * N)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3858: Minimum Bitwise OR From Grid
// https://leetcode.com/problems/minimum-bitwise-or-from-grid/
// Difficulty: Medium
// Time: O(32 * M * N) | Space: O(1)
// Approach: Greedy bit-by-bit from MSB. Try to keep each bit as 0 if
// every row has at least one number with that bit (and all higher
// kept-zero bits) set to 0.

import "fmt"

func MinimumBitwiseOrFromGrid(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	ans := 0
	forbidden := 0

	for bit := 31; bit >= 0; bit-- {
		testForbidden := forbidden | (1 << bit)
		ok := true
		for i := 0; i < m; i++ {
			hasValid := false
			for j := 0; j < n; j++ {
				if grid[i][j]&testForbidden == 0 {
					hasValid = true
					break
				}
			}
			if !hasValid {
				ok = false
				break
			}
		}
		if ok {
			forbidden = testForbidden
		} else {
			ans |= (1 << bit)
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(MinimumBitwiseOrFromGrid([][]int{{1, 5}, {2, 4}})) // Expected: 3

	// Example 2
	fmt.Println(MinimumBitwiseOrFromGrid([][]int{{3, 5}, {6, 4}})) // Expected: 5

	// Example 3
	fmt.Println(MinimumBitwiseOrFromGrid([][]int{{7, 9, 8}})) // Expected: 7
}
```
