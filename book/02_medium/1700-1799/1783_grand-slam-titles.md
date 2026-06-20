# 1783 — Grand Slam Titles

## Deskripsi

**Soal:** [1783. Grand Slam Titles](https://leetcode.com/problems/grand-slam-titles/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func countTitles(players []Player, championships []Championship) map[string]int`

## Solusi Go

```go
package main

// LeetCode #1783: Grand Slam Titles
// https://leetcode.com/problems/grand-slam-titles/
// Difficulty: Medium [Paid] (SQL)
// This is a SQL problem. In Go we simulate the logic.
// Time: O(n), Space: O(n)

import "fmt"

type Player struct {
	ID   int
	Name string
}

type Championship struct {
	Year  int
	WonBy int // player ID
	Tournament string
}

func countTitles(players []Player, championships []Championship) map[string]int {
  // Membuat map untuk pencarian O(1): key → value
	titles := make(map[string]int)
	for _, c := range championships {
		for _, p := range players {
			if p.ID == c.WonBy {
				titles[p.Name]++
				break
			}
		}
	}
	return titles
}

func main() {
	players := []Player{
		{1, "Federer"},
		{2, "Nadal"},
		{3, "Djokovic"},
	}
	champs := []Championship{
		{2023, 3, "Wimbledon"},
		{2023, 3, "US Open"},
		{2023, 2, "French Open"},
	}
	titles := countTitles(players, champs)
	for name, count := range titles {
		fmt.Printf("%s: %d\n", name, count)
	}
}
```
