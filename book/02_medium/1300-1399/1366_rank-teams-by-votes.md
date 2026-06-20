# 1366 — Rank Teams By Votes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func rankTeams(votes []string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n*m + t^2*log(t)) where n = votes, m = teams per vote, t = unique teams  |  **Ruang:** O(t^2) for storing vote positions

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1366: Rank Teams by Votes
// https://leetcode.com/problems/rank-teams-by-votes/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// Test case 1
	fmt.Println(rankTeams([]string{"ABC", "ACB", "ABC", "ACB", "ACB"})) // "ACB"

	// Test case 2
	fmt.Println(rankTeams([]string{"WXYZ", "XYZW"})) // "XWYZ"

	// Test case 3
	fmt.Println(rankTeams([]string{"ZMNAGUEDSJYLBOPHRQICWFXTVK"})) // "ZMNAGUEDSJYLBOPHRQICWFXTVK"

	// Test case 4 - single vote
	fmt.Println(rankTeams([]string{"ABC"})) // "ABC"
}

// Time: O(n*m + t^2*log(t)) where n = votes, m = teams per vote, t = unique teams
// Space: O(t^2) for storing vote positions
func rankTeams(votes []string) string {
	if len(votes) == 0 {
		return ""
	}

	teams := votes[0]
	n := len(teams)

	// Count votes for each position for each team
	// score[team][position] = count
  // HashMap: O(1) lookup
	score := make(map[byte][]int)
	for _, t := range []byte(teams) {
		score[t] = make([]int, n)
	}

	for _, vote := range votes {
		for pos, team := range []byte(vote) {
			score[team][pos]++
		}
	}

	// Sort teams
	teamList := []byte(teams)
  // Custom sort
	sort.Slice(teamList, func(i, j int) bool {
		a, b := teamList[i], teamList[j]
		for pos := 0; pos < n; pos++ {
			if score[a][pos] != score[b][pos] {
				return score[a][pos] > score[b][pos]
			}
		}
		return a < b
	})

	return string(teamList)
}
```
