# 1001 — Grid Illumination

## Deskripsi

**Soal:** [1001. Grid Illumination](https://leetcode.com/problems/grid-illumination/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** HashMap (tabel pencarian O(1))

> **Ide Kunci:** Hash maps for row, col, diagonal, anti-diagonal coverage + lamp set.

## Solusi Go

```go
package main

// LeetCode #1001: Grid Illumination
// https://leetcode.com/problems/grid-illumination/
// Difficulty: Hard
//
// Approach: Hash maps for row, col, diagonal, anti-diagonal coverage + lamp set.
//   - A lamp at (r,c) lights its row r, column c, diagonal r-c, anti-diagonal r+c.
//   - For each query, check if any counter > 0 for the 4 axes of the cell.
//   - Then turn off any lamp in the 3x3 neighborhood (remove its contributions).

import "fmt"

func main() {
	fmt.Println(gridIllumination(5, [][]int{{0, 0}, {4, 4}}, [][]int{{1, 1}, {1, 0}})) // [1,0]
	fmt.Println(gridIllumination(5, [][]int{{0, 0}, {4, 4}}, [][]int{{1, 1}, {1, 1}})) // [1,1]
}

func gridIllumination(n int, lamps [][]int, queries [][]int) []int {
  // Membuat map untuk pencarian O(1): key → value
	rows := make(map[int]int)
  // Membuat map untuk pencarian O(1): key → value
	cols := make(map[int]int)
  // Membuat map untuk pencarian O(1): key → value
	diag := make(map[int]int)
  // Membuat map untuk pencarian O(1): key → value
	anti := make(map[int]int)
  // Membuat map untuk pencarian O(1): key → value
	lampSet := make(map[[2]int]bool)

	for _, l := range lamps {
		r, c := l[0], l[1]
		key := [2]int{r, c}
		if lampSet[key] {
			continue
		}
		lampSet[key] = true
		rows[r]++
		cols[c]++
		diag[r-c]++
		anti[r+c]++
	}

  // Membuat slice untuk menyimpan hasil
	ans := make([]int, len(queries))
	for qi, q := range queries {
		r, c := q[0], q[1]
		if rows[r] > 0 || cols[c] > 0 || diag[r-c] > 0 || anti[r+c] > 0 {
			ans[qi] = 1
		}

		for dr := -1; dr <= 1; dr++ {
			for dc := -1; dc <= 1; dc++ {
				nr, nc := r+dr, c+dc
				if nr < 0 || nr >= n || nc < 0 || nc >= n {
					continue
				}
				key := [2]int{nr, nc}
				if lampSet[key] {
					delete(lampSet, key)
					rows[nr]--
					cols[nc]--
					diag[nr-nc]--
					anti[nr+nc]--
				}
			}
		}
	}
	return ans
}
```
