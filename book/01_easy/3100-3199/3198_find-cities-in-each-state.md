# 3198 — Find Cities In Each State

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func FindCitiesInEachState(data [][]string) [][]string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n). Space: O(n).  |  **Ruang:** O(n).

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3198: Find Cities in Each State
// https://leetcode.com/problems/find-cities-in-each-state/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	cities := [][]string{
		{"New York", "Albany"},
		{"California", "Los Angeles"},
		{"New York", "Buffalo"},
		{"California", "San Francisco"},
		{"Texas", "Houston"},
	}
	result := FindCitiesInEachState(cities)
	for _, r := range result {
		fmt.Println(r)
	}
}

// FindCitiesInEachState groups cities by state and returns them as comma-separated strings.
// Time: O(n log n). Space: O(n).
func FindCitiesInEachState(data [][]string) [][]string {
  // HashMap: O(1) lookup
	stateCities := make(map[string][]string)
	for _, row := range data {
		state, city := row[0], row[1]
		stateCities[state] = append(stateCities[state], city)
	}

	states := make([]string, 0, len(stateCities))
	for s := range stateCities {
		states = append(states, s)
	}
	sort.Strings(states)

  // Matriks 2D
	result := make([][]string, len(states))
	for i, s := range states {
		cities := stateCities[s]
		sort.Strings(cities)
		result[i] = []string{s, strings.Join(cities, ", ")}
	}
	return result
}
```
