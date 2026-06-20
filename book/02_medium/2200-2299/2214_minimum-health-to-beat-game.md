# 2214 — Minimum Health To Beat Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumHealth(damage []int, armor int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2214: Minimum Health to Beat Game
// https://leetcode.com/problems/minimum-health-to-beat-game/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func minimumHealth(damage []int, armor int) int64 {
	var total int64 = 0
	maxDmg := 0
	for _, d := range damage {
		total += int64(d)
		if d > maxDmg {
			maxDmg = d
		}
	}
	// We can use armor to reduce the largest damage
	saved := armor
	if maxDmg < saved {
		saved = maxDmg
	}
	return total - int64(saved) + 1
}

func main() {
	// Test case 1
	fmt.Println(minimumHealth([]int{2, 7, 4, 3}, 4))
	// Expected: 13

	// Test case 2
	fmt.Println(minimumHealth([]int{3, 3, 3}, 0))
	// Expected: 10

	// Test case 3
	fmt.Println(minimumHealth([]int{1, 2, 3, 4}, 5))
	// Expected: 7
}
```
