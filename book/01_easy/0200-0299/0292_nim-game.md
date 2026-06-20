# 0292 — Nim Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func CanWinNim(n int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #292: Nim Game
// https://leetcode.com/problems/nim-game/
// Difficulty: Easy

import "fmt"

// Time: O(1) | Space: O(1)
func CanWinNim(n int) bool {
	return n%4 != 0
}

func main() {
	fmt.Println(CanWinNim(4))
	fmt.Println(CanWinNim(1))
	fmt.Println(CanWinNim(2))
}
```
