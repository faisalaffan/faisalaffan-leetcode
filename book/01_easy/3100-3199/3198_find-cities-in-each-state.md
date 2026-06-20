# 3198 — Find Cities In Each State

## Deskripsi

**Soal:** [3198. Find Cities In Each State](https://leetcode.com/problems/find-cities-in-each-state/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

**Algoritma:** —

## Solusi Go

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
  // Membuat map untuk pencarian O(1): key → value
	stateCities := make(map[string][]string)
	for _, row := range data {
		state, city := row[0], row[1]
		stateCities[state] = append(stateCities[state], city)
	}

  // Membuat slice untuk menyimpan hasil
	states := make([]string, 0, len(stateCities))
	for s := range stateCities {
		states = append(states, s)
	}
	sort.Strings(states)

  // Membuat slice 2D untuk DP/tabel
	result := make([][]string, len(states))
	for i, s := range states {
		cities := stateCities[s]
		sort.Strings(cities)
		result[i] = []string{s, strings.Join(cities, ", ")}
	}
	return result
}
```
