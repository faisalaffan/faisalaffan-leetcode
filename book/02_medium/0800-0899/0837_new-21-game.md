# 0837 — New 21 Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func NewTwoOneGame(n int, k int, maxPts int) float64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Sliding Window, Dynamic Programming

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Sliding Window** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #837: New 21 Game
// https://leetcode.com/problems/new-21-game/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(NewTwoOneGame(10, 1, 10))
	fmt.Println(NewTwoOneGame(6, 1, 10))
	fmt.Println(NewTwoOneGame(21, 17, 10))
}

// Time: O(n) | Space: O(n)
func NewTwoOneGame(n int, k int, maxPts int) float64 {
	if k == 0 || n >= k-1+maxPts {
		return 1.0
	}

	dp := make([]float64, n+1)
	dp[0] = 1.0
	windowSum := 1.0
	var ans float64

	for i := 1; i <= n; i++ {
		dp[i] = windowSum / float64(maxPts)
		if i < k {
			windowSum += dp[i]
		} else {
			ans += dp[i]
		}
		if i >= maxPts {
			windowSum -= dp[i-maxPts]
		}
	}

	return ans
}
```
