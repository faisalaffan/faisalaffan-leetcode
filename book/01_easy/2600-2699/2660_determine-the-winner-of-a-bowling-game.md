# 2660 — Determine The Winner Of A Bowling Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func DetermineTheWinnerOfABowlingGame(player1 []int, player2 []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2660: Determine the Winner of a Bowling Game
// https://leetcode.com/problems/determine-the-winner-of-a-bowling-game/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(DetermineTheWinnerOfABowlingGame([]int{4, 10, 7, 9}, []int{6, 5, 2, 3}))
	fmt.Println(DetermineTheWinnerOfABowlingGame([]int{3, 5, 7, 6}, []int{8, 10, 10, 2}))
}

func DetermineTheWinnerOfABowlingGame(player1 []int, player2 []int) int {
	score1 := computeScore(player1)
	score2 := computeScore(player2)
	if score1 > score2 {
		return 1
	} else if score2 > score1 {
		return 2
	}
	return 0
}

func computeScore(pins []int) int {
	score := 0
	for i, v := range pins {
		score += v
		if (i >= 1 && pins[i-1] == 10) || (i >= 2 && pins[i-2] == 10) {
			score += v
		}
	}
	return score
}
```
