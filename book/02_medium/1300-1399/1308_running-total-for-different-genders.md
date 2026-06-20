# 1308 — Running Total For Different Genders

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func runningTotal(scores []struct { playerName string gender string day string scorePoints int }) []result`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n log n) due to sorting  |  **Ruang:** O(n) for storing results

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1308: Running Total for Different Genders
// https://leetcode.com/problems/running-total-for-different-genders/
// Difficulty: Medium

import "fmt"

func main() {
	scores := []struct {
		playerName  string
		gender      string
		day         string
		scorePoints int
	}{
		{"Aron", "F", "2020-01-01", 17},
		{"Alice", "F", "2020-01-07", 23},
		{"Bajrang", "M", "2020-01-07", 7},
		{"Khali", "M", "2019-12-25", 11},
		{"Slaman", "M", "2019-12-30", 13},
		{"Joe", "M", "2019-12-31", 3},
		{"Jose", "M", "2019-12-18", 2},
		{"Priya", "F", "2019-12-31", 15},
		{"Priyanka", "F", "2019-11-23", 17},
	}

	result := runningTotal(scores)
	for _, r := range result {
		fmt.Printf("%s %s %s %d\n", r.gender, r.day, r.playerName, r.total)
	}
}

type result struct {
	gender     string
	day        string
	playerName string
	total      int
}

// Time: O(n log n) due to sorting
// Space: O(n) for storing results
func runningTotal(scores []struct {
	playerName string
	gender     string
	day        string
	scorePoints int
}) []result {
	// Group by gender, sort each group by day, then gender, then player_name
	// For simplicity, we process in Go using map
	type playerScore struct {
		player, day string
		score       int
	}

  // HashMap: O(1) lookup
	groups := make(map[string][]playerScore)
	for _, s := range scores {
		groups[s.gender] = append(groups[s.gender], playerScore{s.playerName, s.day, s.scorePoints})
	}

	// Sort each group by day, then player_name
	for gender := range groups {
		// Bubble sort for simplicity - in real SQL this is ORDER BY
		group := groups[gender]
  // Linear scan O(n)
		for i := 0; i < len(group); i++ {
			for j := i + 1; j < len(group); j++ {
				if group[j].day < group[i].day ||
					(group[j].day == group[i].day && group[j].player < group[i].player) {
					group[i], group[j] = group[j], group[i]
				}
			}
		}
		groups[gender] = group
	}

	var results []result
	for _, gender := range []string{"F", "M"} {
		running := 0
		for _, ps := range groups[gender] {
			running += ps.score
			results = append(results, result{gender, ps.day, ps.player, running})
		}
	}
	return results
}
```
