# 0486 — Predict The Winner

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func PredictTheWinner(nums []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Dynamic Programming

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n^2)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, n)
  // Range loop: iterasi dengan indeks + nilai
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
