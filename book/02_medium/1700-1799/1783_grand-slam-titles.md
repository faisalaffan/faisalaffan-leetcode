# 1783 — Grand Slam Titles

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func countTitles(players []Player, championships []Championship) map[string]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

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
  // HashMap: O(1) lookup
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
