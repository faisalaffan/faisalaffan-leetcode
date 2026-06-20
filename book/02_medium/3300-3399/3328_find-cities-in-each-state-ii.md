# 3328 — Find Cities In Each State Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func findCityRanking(cities []CityInfo) []CityRank
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(c log c) Space: O(c)  
**Kompleksitas Ruang:** O(c)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Membuat map (HashMap) — pencarian O(1)
	stateCities := make(map[int][]CityInfo)
	for _, c := range cities {
		stateCities[c.StateID] = append(stateCities[c.StateID], c)
	}

	var result []CityRank
	for sid, cs := range stateCities {
  // Custom sort dengan comparator
		sort.Slice(cs, func(i, j int) bool {
			return cs[i].Population > cs[j].Population
		})
		for _, c := range cs {
			result = append(result, CityRank{sid, c.CityName})
		}
	}

  // Custom sort dengan comparator
	sort.Slice(result, func(i, j int) bool {
		if result[i].StateID != result[j].StateID {
			return result[i].StateID < result[j].StateID
		}
		return result[i].CityName < result[j].CityName
	})
	return result
}
```
