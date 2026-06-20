# 0511 — Game Play Analysis I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func GamePlayAnalysisI() string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #511: Game Play Analysis I
// https://leetcode.com/problems/game-play-analysis-i/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func GamePlayAnalysisI() string {
	return "SELECT player_id, MIN(event_date) AS first_login FROM Activity GROUP BY player_id"
}

func main() {
	fmt.Println(GamePlayAnalysisI())
}
```
