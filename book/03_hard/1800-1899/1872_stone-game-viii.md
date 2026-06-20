# 1872 — Stone Game Viii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func stoneGameViii(stones []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1872: Stone Game VIII
// https://leetcode.com/problems/stone-game-viii/
// Difficulty: Hard

import "fmt"

func stoneGameViii(stones []int) int {
	n := len(stones)
  // Alokasi slice integer
	prefix := make([]int, n)
	prefix[0] = stones[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + stones[i]
	}

	// dp[i] = max score difference (current player - opponent) starting from
	// position i. The player may choose any j >= i, j < n-1, take the prefix
	// from position i to j (scoring prefix[j] - base), and leave position j+1
	// for the opponent.
	//
	// Recurrence: dp[i] = max over j >= i of (prefix[j] - base - dp[j+1])
	// where base = 0 (when i=0) or prefix[i-1].
	// This simplifies to: dp[i] = max(prefix[i] - dp[i+1], dp[i+1]).
	//
	// dp[n-1] = 0 (cannot take when only 1 stone remains).

	dp := 0
	for i := n - 2; i >= 0; i-- {
		dp = max(prefix[i]-dp, dp)
	}
	return dp
}

func main() {
	// Test cases
	fmt.Println(stoneGameViii([]int{-1, 2, -3, 4, -5}))
	fmt.Println(stoneGameViii([]int{1, 2, 3, 4, 5}))
}
```
