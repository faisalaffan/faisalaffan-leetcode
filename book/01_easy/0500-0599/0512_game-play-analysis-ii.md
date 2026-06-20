# 0512 — Game Play Analysis Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func GamePlayAnalysisIi() string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #512: Game Play Analysis II
// https://leetcode.com/problems/game-play-analysis-ii/
// Difficulty: Easy [Paid]

import "fmt"

// Time: O(n), Space: O(n)
func GamePlayAnalysisIi() string {
	return "SELECT player_id, device_id FROM Activity WHERE (player_id, event_date) IN (SELECT player_id, MIN(event_date) FROM Activity GROUP BY player_id)"
}

func main() {
	fmt.Println(GamePlayAnalysisIi())
}
```
