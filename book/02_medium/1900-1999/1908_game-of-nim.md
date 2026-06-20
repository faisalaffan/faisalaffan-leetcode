# 1908 — Game Of Nim

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func NimGame(piles []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1908: Game of Nim
// https://leetcode.com/problems/game-of-nim/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(NimGame([]int{1, 2, 3}))
	fmt.Println(NimGame([]int{1, 1, 1}))
	fmt.Println(NimGame([]int{1, 2}))
}

// Time: O(n), Space: O(1)
func NimGame(piles []int) bool {
	xor := 0
	for _, p := range piles {
		xor ^= p
	}
	return xor != 0
}
```
