# 3360 — Stone Removal Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func StoneRemovalGame(n int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(sqrt(n)). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3360: Stone Removal Game
// https://leetcode.com/problems/stone-removal-game/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(StoneRemovalGame(10))
	fmt.Println(StoneRemovalGame(7))
}

// StoneRemovalGame returns true if Alice wins the stone removal game.
// Alice goes first; they remove stones starting from 1 and increase by 1 each turn.
// Alice wins if she can make the last move.
// Time: O(sqrt(n)). Space: O(1).
func StoneRemovalGame(n int) bool {
	turn := 0 // 0 for Alice, 1 for Bob
	remove := 1
	for n >= remove {
		n -= remove
		remove++
		turn = 1 - turn
	}
	// If Alice made the last move, Alice wins, else Bob wins
	return turn != 0
}
```
