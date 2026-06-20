# 1023 — Camelcase Matching

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func camelMatch(queries []string, pattern string) []bool
```

> **💡 Hint:** For each query, use two pointers to match pattern characters

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * (len(query) + len(pattern)))  
**Kompleksitas Ruang:** O(n) for output

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1023: Camelcase Matching
// https://leetcode.com/problems/camelcase-matching/
// Difficulty: Medium
//
// Approach: For each query, use two pointers to match pattern characters
// Time: O(n * (len(query) + len(pattern)))
// Space: O(n) for output

import "fmt"

func main() {
	fmt.Println(camelMatch([]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FB"))
	// [true,false,true,true,false]
	fmt.Println(camelMatch([]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FoBa"))
	// [true,false,true,false,false]
}

func camelMatch(queries []string, pattern string) []bool {
	result := make([]bool, len(queries))

	for i, q := range queries {
		result[i] = matches(q, pattern)
	}

	return result
}

func matches(query, pattern string) bool {
	j := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(query); i++ {
		if j < len(pattern) && query[i] == pattern[j] {
			j++
		} else if query[i] >= 'A' && query[i] <= 'Z' {
			return false
		}
	}
	return j == len(pattern)
}
```
