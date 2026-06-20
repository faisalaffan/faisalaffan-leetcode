# 1690 — Stone Game Vii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func stoneGameVII(stones []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Dynamic Programming, Prefix Sum

**Kompleksitas Waktu:** O(n^2), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1690: Stone Game VII
// https://leetcode.com/problems/stone-game-vii/
// Difficulty: Medium
// Time: O(n^2), Space: O(n)

import "fmt"

func stoneGameVII(stones []int) int {
	n := len(stones)
  // Alokasi slice integer
	prefix := make([]int, n+1)
	for i, v := range stones {
		prefix[i+1] = prefix[i] + v
	}

	// dp[i][j] = max score difference (current player - other player) for subarray i..j
  // Alokasi slice integer
	dp := make([]int, n)

	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			// Remove left: score = sum of rest, then subtract opponent's optimal
			removeLeft := (prefix[j+1] - prefix[i+1]) - dp[i+1]
			// Remove right: score = sum of rest, then subtract opponent's optimal
			removeRight := (prefix[j] - prefix[i]) - dp[i]
			dp[i] = max(removeLeft, removeRight)
		}
	}
	return dp[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(stoneGameVII([]int{5, 3, 1, 4, 2})) // Expected: 6
	fmt.Println(stoneGameVII([]int{7, 90, 5, 1, 100, 10, 10, 2})) // Expected: 122
	fmt.Println(stoneGameVII([]int{1, 2})) // Expected: 2
}
```
