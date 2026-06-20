# 3030 — Find The Grid Of Region Average

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func resultGrid(image [][]int, threshold int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(m*n)  |  **Ruang:** O(m*n)


## 💻 Solusi Go

```go
package main

// LeetCode #3030: Find the Grid of Region Average
// https://leetcode.com/problems/find-the-grid-of-region-average/
// Difficulty: Medium
// Time: O(m*n) | Space: O(m*n)

import "fmt"

func main() {
	fmt.Println(resultGrid([][]int{{5, 6, 7}, {8, 9, 10}, {11, 12, 13}}, 2))
	fmt.Println(resultGrid([][]int{{10, 20, 30}, {15, 25, 30}, {20, 30, 40}}, 5))
	fmt.Println(resultGrid([][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, 0))
}

func resultGrid(image [][]int, threshold int) [][]int {
	m, n := len(image), len(image[0])
  // Matriks 2D
	sum := make([][]int, m)
  // Matriks 2D
	cnt := make([][]int, m)
  // Range loop
	for i := range sum {
		sum[i] = make([]int, n)
		cnt[i] = make([]int, n)
	}

	for i := 0; i+2 < m; i++ {
		for j := 0; j+2 < n; j++ {
			ok := true
			total := 0
		check:
			for x := i; x < i+3; x++ {
				for y := j; y < j+3; y++ {
					total += image[x][y]
					if x > i && abs(image[x][y]-image[x-1][y]) > threshold {
						ok = false
						break check
					}
					if y > j && abs(image[x][y]-image[x][y-1]) > threshold {
						ok = false
						break check
					}
				}
			}
			if !ok {
				continue
			}
			avg := total / 9
			for x := i; x < i+3; x++ {
				for y := j; y < j+3; y++ {
					sum[x][y] += avg
					cnt[x][y]++
				}
			}
		}
	}

  // Matriks 2D
	ans := make([][]int, m)
  // Range loop
	for i := range ans {
		ans[i] = make([]int, n)
		for j := range ans[i] {
			if cnt[i][j] == 0 {
				ans[i][j] = image[i][j]
			} else {
				ans[i][j] = sum[i][j] / cnt[i][j]
			}
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
