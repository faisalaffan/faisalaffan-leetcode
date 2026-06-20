# 3051 — Find Candidates For Data Scientist Position

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func FindCandidatesForDataScientistPosition(candidates [][]string) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

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
  // HashMap: O(1) lookup
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
  // Sort O(n log n)
	sort.Ints(result)
	return result
}

func parseInt(s string) int {
	n := 0
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		n = n*10 + int(s[i]-'0')
	}
	return n
}
```
