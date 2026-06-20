# 0533 — Lonely Pixel Ii

## Deskripsi

**Soal:** [0533. Lonely Pixel Ii](https://leetcode.com/problems/lonely-pixel-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(m * n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #533: Lonely Pixel II
// https://leetcode.com/problems/lonely-pixel-ii/
// Difficulty: Medium [Paid]
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	picture := [][]byte{
		{'W', 'B', 'W', 'B', 'B', 'W'},
		{'W', 'B', 'W', 'B', 'B', 'W'},
		{'W', 'B', 'W', 'B', 'B', 'W'},
		{'W', 'W', 'B', 'W', 'B', 'W'},
	}
	fmt.Println(FindBlackPixel(picture, 3))
}

func FindBlackPixel(picture [][]byte, target int) int {
	m, n := len(picture), len(picture[0])
  // Membuat slice untuk menyimpan hasil
	rows := make([]int, m)
  // Membuat slice untuk menyimpan hasil
	cols := make([]int, n)
  // Membuat map untuk pencarian O(1): key → value
	rowPattern := make(map[string]int)

	for i := 0; i < m; i++ {
		rowStr := ""
		for j := 0; j < n; j++ {
			rowStr += string(picture[i][j])
			if picture[i][j] == 'B' {
				rows[i]++
				cols[j]++
			}
		}
		rowPattern[rowStr]++
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if picture[i][j] == 'B' && rows[i] == target && cols[j] == target {
				// All rows with 'B' in column j must be identical
				allSame := true
				for r := 0; r < m; r++ {
					if picture[r][j] == 'B' && picture[r][j] == picture[i][j] {
						// Check if rows i and r are identical
						for c := 0; c < n; c++ {
							if picture[i][c] != picture[r][c] {
								allSame = false
								break
							}
						}
						if !allSame {
							break
						}
					}
				}
				if allSame {
					count++
				}
			}
		}
	}

	return count
}
```
