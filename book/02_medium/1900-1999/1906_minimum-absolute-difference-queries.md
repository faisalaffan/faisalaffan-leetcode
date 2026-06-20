# 1906 — Minimum Absolute Difference Queries

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func MinDifference(nums []int, queries [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O((n+q)*maxVal), Space: O(n*maxVal)  |  **Ruang:** O(n*maxVal)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1906: Minimum Absolute Difference Queries
// https://leetcode.com/problems/minimum-absolute-difference-queries/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinDifference([]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 3}}))
	fmt.Println(MinDifference([]int{4, 5, 2, 2, 7, 10}, [][]int{{2, 3}, {0, 2}, {0, 5}, {3, 5}}))
}

const maxVal = 100

// Time: O((n+q)*maxVal), Space: O(n*maxVal)
func MinDifference(nums []int, queries [][]int) []int {
	n := len(nums)
	// prefixCount[i][v] = count of value v in nums[0..i-1]
  // Alokasi slice
	prefixCount := make([][maxVal + 1]int, n+1)
	for i := 0; i < n; i++ {
		prefixCount[i+1] = prefixCount[i]
		prefixCount[i+1][nums[i]]++
	}

  // Alokasi slice
	result := make([]int, len(queries))
	for qIdx, q := range queries {
		l, r := q[0], q[1]
		prev := -1
		minDiff := -1
		for v := 1; v <= maxVal; v++ {
			if prefixCount[r+1][v]-prefixCount[l][v] > 0 {
				if prev != -1 {
					diff := v - prev
					if minDiff == -1 || diff < minDiff {
						minDiff = diff
					}
				}
				prev = v
			}
		}
		result[qIdx] = minDiff
	}
	return result
}
```
