# 0512 — Game Play Analysis Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan aturan permainan. Tugasmu menentukan pemenang atau skor optimal.

**Cara berpikir:** DP dari end-state mundur ke awal. Atau analisis pola matematika.

**Fungsi Solusi:** `func GamePlayAnalysisIi() string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)


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
