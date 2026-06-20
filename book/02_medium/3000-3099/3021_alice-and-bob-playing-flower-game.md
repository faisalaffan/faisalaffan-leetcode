# 3021 — Alice And Bob Playing Flower Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func flowerGame(n int, m int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3021: Alice and Bob Playing Flower Game
// https://leetcode.com/problems/alice-and-bob-playing-flower-game/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(flowerGame(3, 2))
	fmt.Println(flowerGame(1, 1))
	fmt.Println(flowerGame(4, 4))
}

func flowerGame(n int, m int) int64 {
	oddN := int64((n + 1) / 2)
	evenN := int64(n / 2)
	oddM := int64((m + 1) / 2)
	evenM := int64(m / 2)
	return oddN*evenM + evenN*oddM
}
```
