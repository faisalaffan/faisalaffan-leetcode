# 0486 — Predict The Winner

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func PredictTheWinner(nums []int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, DP

**Waktu:** O(n^2)  |  **Ruang:** O(n^2)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #486: Predict the Winner
// https://leetcode.com/problems/predict-the-winner/
// Difficulty: Medium
// Time: O(n^2)
// Space: O(n^2)

import "fmt"

func main() {
	fmt.Println(PredictTheWinner([]int{1, 5, 2}))
	fmt.Println(PredictTheWinner([]int{1, 5, 233, 7}))
}

func PredictTheWinner(nums []int) bool {
	n := len(nums)
  // Matriks 2D
	dp := make([][]int, n)
  // Range loop
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}

	for length := 2; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			// Max of (pick left) or (pick right), minus opponent's optimal play
			left := nums[i] - dp[i+1][j]
			right := nums[j] - dp[i][j-1]
			if left > right {
				dp[i][j] = left
			} else {
				dp[i][j] = right
			}
		}
	}

	return dp[0][n-1] >= 0
}
```
