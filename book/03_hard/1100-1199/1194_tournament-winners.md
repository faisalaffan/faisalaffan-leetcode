# 1194 — Tournament Winners

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func getTournamentWinners(players []Player, matches []Match) []GroupWinner
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1194: Tournament Winners
// https://leetcode.com/problems/tournament-winners/
// Difficulty: Hard [Paid]
//
// Given players (with group_id) and matches (first_player, second_player,
// first_score, second_score), find the winner of each group — the player
// with the highest total score in their group. Ties are broken by the lowest
// player_id.

import (
	"fmt"
	"sort"
)

// Player represents a tournament player.
type Player struct {
	PlayerID    int
	GroupID     int
}

// Match represents a played match.
type Match struct {
	MatchID     int
	FirstPlayer int
	FirstScore  int
	SecondPlayer int
	SecondScore int
}

// GroupWinner is the output format.
type GroupWinner struct {
	GroupID  int
	PlayerID int
}

func main() {
	players := []Player{
		{15, 1}, {25, 1}, {30, 1},
		{45, 2}, {10, 2}, {35, 2},
		{50, 3}, {20, 3}, {40, 3},
	}

	matches := []Match{
		{1, 15, 2, 25, 1},
		{2, 30, 3, 15, 1},
		{3, 10, 4, 35, 0},
		{4, 45, 2, 10, 3},
		{5, 20, 1, 40, 5},
		{6, 50, 2, 20, 3},
	}

	winners := getTournamentWinners(players, matches)
	for _, w := range winners {
		fmt.Printf("group=%d winner=%d\n", w.GroupID, w.PlayerID)
	}
}

// getTournamentWinners determines each group's winner.
func getTournamentWinners(players []Player, matches []Match) []GroupWinner {
	// Player -> group mapping
  // Membuat map (HashMap) — pencarian O(1)
	playerGroup := make(map[int]int)
  // Membuat map (HashMap) — pencarian O(1)
	groupPlayers := make(map[int][]int) // group -> player list

	for _, p := range players {
		playerGroup[p.PlayerID] = p.GroupID
		groupPlayers[p.GroupID] = append(groupPlayers[p.GroupID], p.PlayerID)
	}

	// Compute total score per player
  // Membuat map (HashMap) — pencarian O(1)
	scores := make(map[int]int)
	for _, m := range matches {
		scores[m.FirstPlayer] += m.FirstScore
		scores[m.SecondPlayer] += m.SecondScore
	}

	// Find winner per group
	var winners []GroupWinner
  // Alokasi slice integer
	groupIDs := make([]int, 0, len(groupPlayers))
	for g := range groupPlayers {
		groupIDs = append(groupIDs, g)
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(groupIDs)

	for _, gid := range groupIDs {
		bestPlayer := -1
		bestScore := -1

		for _, pid := range groupPlayers[gid] {
			score := scores[pid]
			if score > bestScore || (score == bestScore && (bestPlayer == -1 || pid < bestPlayer)) {
				bestScore = score
				bestPlayer = pid
			}
		}

		winners = append(winners, GroupWinner{GroupID: gid, PlayerID: bestPlayer})
	}

	return winners
}
```
