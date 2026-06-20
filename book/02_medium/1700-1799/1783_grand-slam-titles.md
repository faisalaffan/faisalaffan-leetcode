# 1783 — Grand Slam Titles

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func countTitles(players []Player, championships []Championship) map[string]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Membuat map (HashMap) — pencarian O(1)
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
