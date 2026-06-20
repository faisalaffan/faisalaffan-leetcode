# 0835 — Image Overlap

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func ImageOverlap(img1 [][]int, img2 [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^4)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #835: Image Overlap
// https://leetcode.com/problems/image-overlap/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(ImageOverlap([][]int{{1, 1, 0}, {0, 1, 0}, {0, 1, 0}}, [][]int{{0, 0, 0}, {0, 1, 1}, {0, 0, 1}}))
	fmt.Println(ImageOverlap([][]int{{1}}, [][]int{{1}}))
	fmt.Println(ImageOverlap([][]int{{0}}, [][]int{{0}}))
}

// Time: O(n^4) | Space: O(1)
func ImageOverlap(img1 [][]int, img2 [][]int) int {
	n := len(img1)
	ans := 0

	for a := 1 - n; a < n; a++ {
		for b := 1 - n; b < n; b++ {
			count := 0
			for i := max(a, 0); i < min(n, n+a); i++ {
				for j := max(b, 0); j < min(n, n+b); j++ {
					if img2[i][j] == 1 && img1[i-a][j-b] == 1 {
						count++
					}
				}
			}
			if count > ans {
				ans = count
			}
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
