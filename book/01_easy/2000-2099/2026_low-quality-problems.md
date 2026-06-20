# 2026 — Low Quality Problems

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func LowQualityProblems(problems [][3]string) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2026: Low-Quality Problems
// https://leetcode.com/problems/low-quality-problems/
// Difficulty: Easy [Paid] (SQL)

import (
	"fmt"
	"sort"
)

func main() {
	// Problems: (problem_id, likes, dislikes)
	problems := [][3]string{{"1", "10", "5"}, {"2", "3", "10"}, {"3", "2", "2"}}
	fmt.Println(LowQualityProblems(problems)) // [2]
}

// Time: O(n log n), Space: O(n)
func LowQualityProblems(problems [][3]string) []int {
	var result []int
	for _, p := range problems {
		id := 0
		for _, c := range p[0] {
			id = id*10 + int(c-'0')
		}
		likes := 0
		for _, c := range p[1] {
			likes = likes*10 + int(c-'0')
		}
		dislikes := 0
		for _, c := range p[2] {
			dislikes = dislikes*10 + int(c-'0')
		}

		// Low-quality: likes / (likes + dislikes) < 0.6
		total := likes + dislikes
		if total > 0 && likes*5 < total*3 {
			result = append(result, id)
		}
	}
  // Sort O(n log n)
	sort.Ints(result)
	return result
}
```
