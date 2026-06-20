# 1510 — Stone Game Iv

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func winnerSquareGame(n int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1510: Stone Game IV
// https://leetcode.com/problems/stone-game-iv/
// Difficulty: Hard
//
// Alice and Bob take turns removing stones. In one move, a player can remove
// k*k stones (a perfect square). Alice goes first. Return true if Alice can win
// (playing optimally), false otherwise.

import (
	"fmt"
	"math"
)

// winnerSquareGame returns true if Alice can win the stone game.
func winnerSquareGame(n int) bool {
	// dp[i] = true if the current player can win with i stones remaining
	dp := make([]bool, n+1)

	for i := 1; i <= n; i++ {
		maxSquare := int(math.Sqrt(float64(i)))
		for k := 1; k <= maxSquare; k++ {
			sq := k * k
			// If there's a move that leaves the opponent in a losing position
			if !dp[i-sq] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}

func main() {
	// Test case 1
	n1 := 1
	result1 := winnerSquareGame(n1)
	fmt.Printf("Test 1: n=%d => %v (expected true, Alice removes 1)\n", n1, result1)

	// Test case 2
	n2 := 2
	result2 := winnerSquareGame(n2)
	fmt.Printf("Test 2: n=%d => %v (expected false, Alice removes 1, Bob removes 1)\n", n2, result2)

	// Test case 3
	n3 := 4
	result3 := winnerSquareGame(n3)
	fmt.Printf("Test 3: n=%d => %v (expected true, Alice removes 4)\n", n3, result3)

	// Test case 4
	n4 := 7
	result4 := winnerSquareGame(n4)
	fmt.Printf("Test 4: n=%d => %v (expected false)\n", n4, result4)

	// Test case 5
	n5 := 17
	result5 := winnerSquareGame(n5)
	fmt.Printf("Test 5: n=%d => %v (expected false)\n", n5, result5)

	// Test case 6: larger number
	n6 := 100
	result6 := winnerSquareGame(n6)
	fmt.Printf("Test 6: n=%d => %v\n", n6, result6)

	// Test case 7
	fmt.Println("\nFirst 20 results:")
	for n := 1; n <= 20; n++ {
		fmt.Printf("n=%d: %v\n", n, winnerSquareGame(n))
	}
}
```
