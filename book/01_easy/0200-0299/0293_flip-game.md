# 0293 — Flip Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func GeneratePossibleNextMoves(currentState string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n) for output

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #293: Flip Game
// https://leetcode.com/problems/flip-game/
// Difficulty: Easy [Paid]

import "fmt"

// Time: O(n) | Space: O(n) for output
func GeneratePossibleNextMoves(currentState string) []string {
	var res []string
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(currentState)-1; i++ {
		if currentState[i] == '+' && currentState[i+1] == '+' {
			flipped := currentState[:i] + "--" + currentState[i+2:]
			res = append(res, flipped)
		}
	}
	return res
}

func main() {
	fmt.Println(GeneratePossibleNextMoves("++++"))
	fmt.Println(GeneratePossibleNextMoves("+"))
}
```
