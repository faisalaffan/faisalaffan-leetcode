# 3198 — Find Cities In Each State

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindCitiesInEachState(data [][]string) [][]string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Membuat map (HashMap) — pencarian O(1)
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

  // Membuat matriks/slice 2D untuk DP
	result := make([][]string, len(states))
	for i, s := range states {
		cities := stateCities[s]
		sort.Strings(cities)
		result[i] = []string{s, strings.Join(cities, ", ")}
	}
	return result
}
```
