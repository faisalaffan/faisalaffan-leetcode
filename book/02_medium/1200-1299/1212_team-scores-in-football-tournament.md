# 1212 — Team Scores In Football Tournament

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func calculateTeamScores(teams []string, results []matchResult) []teamScore
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1212: Team Scores in Football Tournament
// https://leetcode.com/problems/team-scores-in-football-tournament/
// Difficulty: Medium [Paid]

// Teams get 3 pts for win, 1 for draw, 0 for loss.
// Calculate total points per team, sorted by points desc then name asc.

// Time: O(n log n)
// Space: O(n)

type matchResult struct {
	host       string
	guest      string
	hostScore  int
	guestScore int
}

type teamScore struct {
	name   string
	points int
}

func calculateTeamScores(teams []string, results []matchResult) []teamScore {
  // Membuat map (HashMap) — pencarian O(1)
	points := make(map[string]int)

	for _, r := range results {
		if r.hostScore > r.guestScore {
			points[r.host] += 3
		} else if r.hostScore < r.guestScore {
			points[r.guest] += 3
		} else {
			points[r.host] += 1
			points[r.guest] += 1
		}
	}

	scores := make([]teamScore, 0, len(teams))
	for _, t := range teams {
		scores = append(scores, teamScore{t, points[t]})
	}

  // Custom sort dengan comparator
	sort.Slice(scores, func(i, j int) bool {
		if scores[i].points != scores[j].points {
			return scores[i].points > scores[j].points
		}
		return scores[i].name < scores[j].name
	})

	return scores
}

func main() {
	results := []matchResult{
		{"A", "B", 2, 1},
		{"C", "A", 1, 1},
		{"B", "C", 0, 3},
	}
	scores := calculateTeamScores([]string{"A", "B", "C"}, results)
	for _, s := range scores {
		fmt.Printf("%s: %d\n", s.name, s.points)
	}
}
```
