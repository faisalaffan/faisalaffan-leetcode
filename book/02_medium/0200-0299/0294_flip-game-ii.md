# 0294 — Flip Game Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func canWin(currentState string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n!!) worst case with memo, Space: O(n!)  
**Kompleksitas Ruang:** O(n!)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #294: Flip Game II
// https://leetcode.com/problems/flip-game-ii/
// Difficulty: Medium [Paid]
// Time: O(n!!) worst case with memo, Space: O(n!)

import "fmt"

func canWin(currentState string) bool {
  // Membuat map (HashMap) — pencarian O(1)
	memo := make(map[string]bool)
	return canWinHelper(currentState, memo)
}

func canWinHelper(state string, memo map[string]bool) bool {
	if res, ok := memo[state]; ok {
		return res
	}

	bytes := []byte(state)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(state)-1; i++ {
		if bytes[i] == '+' && bytes[i+1] == '+' {
			bytes[i], bytes[i+1] = '-', '-'
			opponentWins := canWinHelper(string(bytes), memo)
			bytes[i], bytes[i+1] = '+', '+'

			if !opponentWins {
				memo[state] = true
				return true
			}
		}
	}

	memo[state] = false
	return false
}

func main() {
	fmt.Println(canWin("++++"))
	fmt.Println(canWin("+"))
	fmt.Println(canWin("+++"))
}
```
