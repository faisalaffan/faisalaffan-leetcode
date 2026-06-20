# 2026 — Low Quality Problems

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func LowQualityProblems(problems [][3]string) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Urutkan secara ascending — O(n log n)
	sort.Ints(result)
	return result
}
```
