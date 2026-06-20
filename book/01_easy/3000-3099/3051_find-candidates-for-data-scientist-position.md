# 3051 — Find Candidates For Data Scientist Position

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindCandidatesForDataScientistPosition(candidates [][]string) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3051: Find Candidates for Data Scientist Position
// https://leetcode.com/problems/find-candidates-for-data-scientist-position/
// Difficulty: Easy [Paid - Database]
//
// Note: This is a SQL problem on LeetCode. In Go, we implement the
// equivalent logic: find candidates with Python, Tableau, and PostgreSQL skills.

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: findCandidates
	// Input: [candidate_id, skill]
	candidates := [][]string{
		{"101", "Python"},
		{"101", "Tableau"},
		{"101", "PostgreSQL"},
		{"102", "Python"},
		{"102", "Tableau"},
		{"103", "Python"},
		{"103", "PostgreSQL"},
	}
	fmt.Println(FindCandidatesForDataScientistPosition(candidates))
	// [101]
}

// Time: O(n log n) | Space: O(n)
// LeetCode submission name: findCandidates
func FindCandidatesForDataScientistPosition(candidates [][]string) []int {
  // Membuat map (HashMap) — pencarian O(1)
	skills := make(map[int]map[string]bool)

	for _, row := range candidates {
		id := parseInt(row[0])
		skill := row[1]
		if skills[id] == nil {
			skills[id] = make(map[string]bool)
		}
		skills[id][skill] = true
	}

	result := []int{}
	for id, s := range skills {
		if s["Python"] && s["Tableau"] && s["PostgreSQL"] {
			result = append(result, id)
		}
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(result)
	return result
}

func parseInt(s string) int {
	n := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		n = n*10 + int(s[i]-'0')
	}
	return n
}
```
