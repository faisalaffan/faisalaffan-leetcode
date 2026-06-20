# 1039 — Minimum Score Triangulation Of Polygon

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minScoreTriangulation(values []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(n^3)  |  **Ruang:** O(n^2)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1039: Minimum Score Triangulation of Polygon
// https://leetcode.com/problems/minimum-score-triangulation-of-polygon/
// Difficulty: Medium
//
// Approach: Interval DP. dp[i][j] = min score triangulating polygon from i to j.
// Time: O(n^3)
// Space: O(n^2)

import "fmt"

func main() {
	fmt.Println(minScoreTriangulation([]int{1, 2, 3}))    // 6
	fmt.Println(minScoreTriangulation([]int{3, 7, 4, 5})) // 144
}

func minScoreTriangulation(values []int) int {
	n := len(values)
  // Matriks 2D
	dp := make([][]int, n)
  // Range loop
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for length := 2; length < n; length++ {
		for i := 0; i+length < n; i++ {
			j := i + length
			dp[i][j] = 1<<31 - 1
			for k := i + 1; k < j; k++ {
				score := dp[i][k] + dp[k][j] + values[i]*values[j]*values[k]
				if score < dp[i][j] {
					dp[i][j] = score
				}
			}
		}
	}

	return dp[0][n-1]
}
```
