# 1406 — Stone Game Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func stoneGameIII(stoneValue []int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1406: Stone Game III
// https://leetcode.com/problems/stone-game-iii/
// Difficulty: Hard

import "fmt"

func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)
  // Alokasi slice integer
	dp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		dp[i] = stoneValue[i] - dp[i+1]
		if i+2 <= n {
			if sum := stoneValue[i] + stoneValue[i+1] - dp[i+2]; sum > dp[i] {
				dp[i] = sum
			}
		}
		if i+3 <= n {
			if sum := stoneValue[i] + stoneValue[i+1] + stoneValue[i+2] - dp[i+3]; sum > dp[i] {
				dp[i] = sum
			}
		}
	}
	if dp[0] > 0 {
		return "Alice"
	} else if dp[0] < 0 {
		return "Bob"
	}
	return "Tie"
}

func main() {
	// Example: [1,2,3,7] -> "Bob"
	fmt.Println(stoneGameIII([]int{1, 2, 3, 7}))
}
```
