# 0390 — Elimination Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func lastRemaining(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #390: Elimination Game
// https://leetcode.com/problems/elimination-game/
// Difficulty: Medium
// Time: O(log n) | Space: O(1)

import "fmt"

func lastRemaining(n int) int {
	head := 1
	remaining := n
	step := 1
	leftToRight := true

	for remaining > 1 {
		if leftToRight || remaining%2 == 1 {
			head += step
		}
		remaining /= 2
		step *= 2
		leftToRight = !leftToRight
	}
	return head
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", lastRemaining(9))
	// Expected: 6

	// Test case 2
	fmt.Println("Test 2:", lastRemaining(1))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", lastRemaining(100))
	// Expected: 54
}
```
