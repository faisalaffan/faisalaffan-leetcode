# 0534 — Game Play Analysis Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func PlayerActivityReport(activities [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n) for sorting  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #534: Game Play Analysis III
// https://leetcode.com/problems/game-play-analysis-iii/
// Difficulty: Medium [Paid]
// Time: O(n log n) for sorting
// Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	// Sample player activity data
	activities := [][]int{
		{1, 2016, 5},  // player_id=1, event_date=2016, games_played=5
		{1, 2017, 10}, // player_id=1, event_date=2017, games_played=10
		{2, 2016, 15}, // player_id=2, event_date=2016, games_played=15
	}
	fmt.Println(PlayerActivityReport(activities))
}

func PlayerActivityReport(activities [][]int) [][]int {
  // Custom sort dengan comparator
	sort.Slice(activities, func(i, j int) bool {
		if activities[i][0] != activities[j][0] {
			return activities[i][0] < activities[j][0]
		}
		return activities[i][1] < activities[j][1]
	})

	result := [][]int{}
	runningSum := 0
	for i, act := range activities {
		if i > 0 && act[0] != activities[i-1][0] {
			runningSum = 0
		}
		runningSum += act[2]
		result = append(result, []int{act[0], act[1], runningSum})
	}
	return result
}
```
