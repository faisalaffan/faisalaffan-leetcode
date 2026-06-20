# 0511 — Game Play Analysis I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan aturan permainan. Tugasmu menentukan pemenang atau skor optimal.

**Cara berpikir:** DP dari end-state mundur ke awal. Atau analisis pola matematika.

**Fungsi Solusi:** `func GamePlayAnalysisI() string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)


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
