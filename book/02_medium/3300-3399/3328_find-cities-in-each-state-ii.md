# 3328 — Find Cities In Each State Ii

## Deskripsi

**Soal:** [3328. Find Cities In Each State Ii](https://leetcode.com/problems/find-cities-in-each-state-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(c log c) Space: O(c)  
**Kompleksitas Ruang:** O(c)

**Algoritma:** —

## Solusi Go

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
  // Membuat map untuk pencarian O(1): key → value
	stateCities := make(map[int][]CityInfo)
	for _, c := range cities {
		stateCities[c.StateID] = append(stateCities[c.StateID], c)
	}

	var result []CityRank
	for sid, cs := range stateCities {
		sort.Slice(cs, func(i, j int) bool {
			return cs[i].Population > cs[j].Population
		})
		for _, c := range cs {
			result = append(result, CityRank{sid, c.CityName})
		}
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].StateID != result[j].StateID {
			return result[i].StateID < result[j].StateID
		}
		return result[i].CityName < result[j].CityName
	})
	return result
}
```
