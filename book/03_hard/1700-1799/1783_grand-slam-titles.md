# 1783 — Grand Slam Titles

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func grandSlamTitles(players []Player, championships []Championship) []PlayerGrandSlams`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1783: Grand Slam Titles
// https://leetcode.com/problems/grand-slam-titles/
// Difficulty: Hard [Premium]

import (
	"fmt"
	"sort"
)

type Player struct {
	PlayerId int
	Name     string
}

type Championship struct {
	Year      int
	Wimbledon int
	FrOpen    int
	UsOpen    int
	AuOpen    int
}

type PlayerGrandSlams struct {
	PlayerId        int
	Name            string
	GrandSlamsCount int
}

func main() {
	players := []Player{
		{1, "Nadal"},
		{2, "Federer"},
		{3, "Novak"},
	}

	championships := []Championship{
		{2018, 1, 1, 1, 1},
		{2019, 1, 1, 2, 2},
		{2020, 2, 1, 1, 1},
	}

	result := grandSlamTitles(players, championships)
	fmt.Println("Test 1 - Grand Slam Titles:")
	for _, r := range result {
		fmt.Printf("  PlayerId: %d, Name: %s, GrandSlamsCount: %d\n",
			r.PlayerId, r.Name, r.GrandSlamsCount)
	}

	// Test 2: no wins
	players2 := []Player{
		{1, "Player1"},
		{2, "Player2"},
	}
	championships2 := []Championship{
		{2020, 3, 3, 3, 3},
	}
	result2 := grandSlamTitles(players2, championships2)
	fmt.Println("\nTest 2 - Player 3 won but not in players list:")
	for _, r := range result2 {
		fmt.Printf("  PlayerId: %d, Name: %s, Count: %d\n", r.PlayerId, r.Name, r.GrandSlamsCount)
	}
}

func grandSlamTitles(players []Player, championships []Championship) []PlayerGrandSlams {
  // HashMap: O(1) lookup
	counts := make(map[int]int)
	for _, c := range championships {
		counts[c.Wimbledon]++
		counts[c.FrOpen]++
		counts[c.UsOpen]++
		counts[c.AuOpen]++
	}

  // HashMap: O(1) lookup
	playerMap := make(map[int]string)
	for _, p := range players {
		playerMap[p.PlayerId] = p.Name
	}

	result := make([]PlayerGrandSlams, 0)
	for _, p := range players {
		if cnt, ok := counts[p.PlayerId]; ok && cnt > 0 {
			result = append(result, PlayerGrandSlams{
				PlayerId:        p.PlayerId,
				Name:            p.Name,
				GrandSlamsCount: cnt,
			})
		}
	}

  // Custom sort
	sort.Slice(result, func(i, j int) bool {
		if result[i].GrandSlamsCount != result[j].GrandSlamsCount {
			return result[i].GrandSlamsCount > result[j].GrandSlamsCount
		}
		return result[i].PlayerId < result[j].PlayerId
	})

	return result
}
```
