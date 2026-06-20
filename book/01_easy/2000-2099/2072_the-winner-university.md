# 2072 — The Winner University

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func TheWinnerUniversity(newYork []int, california []int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2072: The Winner University
// https://leetcode.com/problems/the-winner-university/
// Difficulty: Easy [Paid] (SQL)

import "fmt"

func main() {
	// New York students: (score)
	newYork := []int{90, 85, 78}
	// California students: (score)
	california := []int{92, 88, 70}
	fmt.Println(TheWinnerUniversity(newYork, california)) // "California University"
}

// Time: O(n), Space: O(1)
func TheWinnerUniversity(newYork []int, california []int) string {
	nyScore := 0
	for _, s := range newYork {
		nyScore += s
	}
	caScore := 0
	for _, s := range california {
		caScore += s
	}

	if nyScore > caScore {
		return "New York University"
	} else if caScore > nyScore {
		return "California University"
	}
	return "No Winner"
}
```
