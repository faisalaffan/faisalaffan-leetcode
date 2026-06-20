# 3252 — Premier League Table Ranking Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func premierLeagueRanking(teams [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3252: Premier League Table Ranking II
// https://leetcode.com/problems/premier-league-table-ranking-ii/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type team struct {
	id     int
	points int
	gDiff  int
}

func premierLeagueRanking(teams [][]int) []int {
	var list []team
	for _, t := range teams {
		list = append(list, team{t[0], t[1], t[2]})
	}

  // Custom sort dengan comparator
	sort.Slice(list, func(i, j int) bool {
		if list[i].points != list[j].points {
			return list[i].points > list[j].points
		}
		if list[i].gDiff != list[j].gDiff {
			return list[i].gDiff > list[j].gDiff
		}
		return list[i].id < list[j].id
	})

  // Alokasi slice integer
	ans := make([]int, len(list))
	for i, t := range list {
		ans[i] = t.id
	}
	return ans
}

func main() {
	fmt.Println(premierLeagueRanking([][]int{{1, 10, 5}, {2, 8, 8}, {3, 10, 3}})) // Expected: [1 3 2]
	fmt.Println(premierLeagueRanking([][]int{{1, 6, 2}, {2, 6, 2}}))              // Expected: [1 2]
}
```
