# 1025 — Divisor Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func divisorGame(n int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1025: Divisor Game
// https://leetcode.com/problems/divisor-game/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(divisorGame(2)) // true
	fmt.Println(divisorGame(3)) // false
	fmt.Println(divisorGame(4)) // true
}

// LeetCode submission: divisorGame
func divisorGame(n int) bool {
	return n%2 == 0
}
```
