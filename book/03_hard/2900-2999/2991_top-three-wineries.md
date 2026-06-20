# 2991 — Top Three Wineries

## Deskripsi

**Soal:** [2991. Top Three Wineries](https://leetcode.com/problems/top-three-wineries/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Trie (pohon awalan)

**Fungsi Solusi:** `func topThreeWineries(data []Winery) []string`

## Solusi Go

```go
package main

// LeetCode #2991: Top Three Wineries (SQL simulation)
// https://leetcode.com/problems/top-three-wineries/
// Difficulty: Hard [Paid]
//
// For each country, find the top 3 wineries based on total points.
// Sort wineries by total points descending, then name ascending.
// If fewer than 3 wineries, fill with "No <rank> winery".

import (
	"fmt"
	"sort"
)

type Winery struct {
	Country string
	Winery  string
	Points  int
}

func topThreeWineries(data []Winery) []string {
	type key struct{ country, winery string }
  // Membuat map untuk pencarian O(1): key → value
	totals := make(map[key]int)
	for _, w := range data {
		k := key{w.Country, w.Winery}
		totals[k] += w.Points
	}

	type wineryScore struct {
		name   string
		points int
	}
  // Membuat map untuk pencarian O(1): key → value
	byCountry := make(map[string][]wineryScore)
	for k, pts := range totals {
		byCountry[k.country] = append(byCountry[k.country], wineryScore{k.winery, pts})
	}

  // Membuat slice untuk menyimpan hasil
	countries := make([]string, 0, len(byCountry))
	for c := range byCountry {
		countries = append(countries, c)
	}
	sort.Strings(countries)

	var result []string
	for _, c := range countries {
		list := byCountry[c]
		sort.Slice(list, func(i, j int) bool {
			if list[i].points != list[j].points {
				return list[i].points > list[j].points
			}
			return list[i].name < list[j].name
		})
		row := fmt.Sprintf("%s|%s (%d)", c, list[0].name, list[0].points)
		if len(list) >= 2 {
			row += fmt.Sprintf("|%s (%d)", list[1].name, list[1].points)
		} else {
			row += "|No second winery"
		}
		if len(list) >= 3 {
			row += fmt.Sprintf("|%s (%d)", list[2].name, list[2].points)
		} else {
			row += "|No third winery"
		}
		result = append(result, row)
	}
	return result
}

func main() {
	// Example from problem description
	data := []Winery{
		{"USA", "RoyalVines", 47},
		{"USA", "RoyalVines", 39},
		{"USA", "SunsetCellars", 85},
		{"USA", "HarmonyHill", 100},
		{"France", "Bordeaux", 95},
		{"France", "Bordeaux", 3},
		{"France", "Loire", 88},
	}
	fmt.Println("Test 1:")
	for _, r := range topThreeWineries(data) {
		fmt.Println(r)
	}

	// Single winery per country
	fmt.Println("\nTest 2 (single winery per country):")
	data2 := []Winery{
		{"Italy", "Chianti", 90},
		{"Italy", "Chianti", 80},
		{"Spain", "Rioja", 85},
	}
	for _, r := range topThreeWineries(data2) {
		fmt.Println(r)
	}

	// Multiple countries with ties
	fmt.Println("\nTest 3 (ties):")
	data3 := []Winery{
		{"A", "X", 100},
		{"A", "Y", 100},
		{"A", "Z", 50},
		{"B", "P", 200},
	}
	for _, r := range topThreeWineries(data3) {
		fmt.Println(r)
	}
}
```
