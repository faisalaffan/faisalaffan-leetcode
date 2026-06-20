# 0877 — Stone Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func StoneGame(piles []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #877: Stone Game
// https://leetcode.com/problems/stone-game/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(StoneGame([]int{5, 3, 4, 5}))
	fmt.Println(StoneGame([]int{3, 7, 2, 5}))
	fmt.Println(StoneGame([]int{1, 100, 3, 2}))
}

// Time: O(1) | Space: O(1)
// Alex always wins because there are an even number of piles
// and total stones is odd (no ties), with Alex going first.
func StoneGame(piles []int) bool {
	return true
}
```
