# 3328 — Find Cities In Each State Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func findCityRanking(cities []CityInfo) []CityRank`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(c log c) Space: O(c)  |  **Ruang:** O(c)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3328: Find Cities in Each State II
// https://leetcode.com/problems/find-cities-in-each-state-ii/
// Difficulty: Medium
// Time: O(c log c) Space: O(c)

import (
	"fmt"
	"sort"
)

func main() {
	cities := []CityInfo{
		{1, "NY", 1000000},
		{1, "LA", 500000},
		{2, "Chicago", 800000},
		{2, "Houston", 600000},
		{2, "Phoenix", 400000},
	}
	fmt.Println(findCityRanking(cities))
}

type CityInfo struct {
	StateID    int
	CityName   string
	Population int
}

type CityRank struct {
	StateID int
	CityName string
}

func findCityRanking(cities []CityInfo) []CityRank {
  // HashMap: O(1) lookup
	stateCities := make(map[int][]CityInfo)
	for _, c := range cities {
		stateCities[c.StateID] = append(stateCities[c.StateID], c)
	}

	var result []CityRank
	for sid, cs := range stateCities {
  // Custom sort
		sort.Slice(cs, func(i, j int) bool {
			return cs[i].Population > cs[j].Population
		})
		for _, c := range cs {
			result = append(result, CityRank{sid, c.CityName})
		}
	}

  // Custom sort
	sort.Slice(result, func(i, j int) bool {
		if result[i].StateID != result[j].StateID {
			return result[i].StateID < result[j].StateID
		}
		return result[i].CityName < result[j].CityName
	})
	return result
}
```
