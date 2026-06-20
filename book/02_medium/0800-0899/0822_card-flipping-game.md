# 0822 — Card Flipping Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func CardFlippingGame(fronts []int, backs []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #822: Card Flipping Game
// https://leetcode.com/problems/card-flipping-game/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CardFlippingGame([]int{1, 2, 4, 4, 7}, []int{1, 3, 4, 1, 3}))
	fmt.Println(CardFlippingGame([]int{1, 1}, []int{1, 2}))
	fmt.Println(CardFlippingGame([]int{1, 1}, []int{2, 2}))
}

// Time: O(n) | Space: O(n)
func CardFlippingGame(fronts []int, backs []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	blocked := make(map[int]bool)
  // Range loop: iterasi dengan indeks + nilai
	for i := range fronts {
		if fronts[i] == backs[i] {
			blocked[fronts[i]] = true
		}
	}

	ans := 2001
	for _, v := range fronts {
		if !blocked[v] && v < ans {
			ans = v
		}
	}
	for _, v := range backs {
		if !blocked[v] && v < ans {
			ans = v
		}
	}

	if ans == 2001 {
		return 0
	}
	return ans
}
```
