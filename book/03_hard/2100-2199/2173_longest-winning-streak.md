# 2173 — Longest Winning Streak

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func longestWinningStreak(matches []match) []playerStreak
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2173: Longest Winning Streak
// https://leetcode.com/problems/longest-winning-streak/
// Difficulty: Hard [Paid]
//
// Given match results (player_id, match_date, result), find each player's
// longest consecutive winning streak. Sort matches by player then date,
// then scan for consecutive wins.

import (
	"fmt"
	"sort"
)

type match struct {
	playerID int
	date     int // represented as integer for simplicity
	result   string
}

type playerStreak struct {
	PlayerID       int
	LongestStreak int
}

func main() {
	// Example
	matches1 := []match{
		{1, 1, "Win"},
		{1, 2, "Win"},
		{1, 3, "Loss"},
		{1, 4, "Win"},
		{2, 1, "Win"},
		{2, 3, "Loss"},
		{2, 5, "Win"},
	}
	fmt.Println(longestWinningStreak(matches1))

	// All wins
	matches2 := []match{
		{1, 1, "Win"},
		{1, 2, "Win"},
		{1, 3, "Win"},
	}
	fmt.Println(longestWinningStreak(matches2))

	// No wins
	matches3 := []match{
		{1, 1, "Loss"},
		{1, 2, "Loss"},
	}
	fmt.Println(longestWinningStreak(matches3))

	// Single match
	matches4 := []match{
		{1, 1, "Win"},
	}
	fmt.Println(longestWinningStreak(matches4))
}

func longestWinningStreak(matches []match) []playerStreak {
	if len(matches) == 0 {
		return nil
	}

	// Sort by playerID, then by date
  // Custom sort dengan comparator
	sort.Slice(matches, func(i, j int) bool {
		if matches[i].playerID != matches[j].playerID {
			return matches[i].playerID < matches[j].playerID
		}
		return matches[i].date < matches[j].date
	})

	var result []playerStreak
	currentPlayer := matches[0].playerID
	currentStreak := 0
	maxStreak := 0

	for _, m := range matches {
		if m.playerID != currentPlayer {
			result = append(result, playerStreak{currentPlayer, maxStreak})
			currentPlayer = m.playerID
			currentStreak = 0
			maxStreak = 0
		}

		if m.result == "Win" {
			currentStreak++
			if currentStreak > maxStreak {
				maxStreak = currentStreak
			}
		} else {
			currentStreak = 0
		}
	}

	result = append(result, playerStreak{currentPlayer, maxStreak})
	return result
}
```
