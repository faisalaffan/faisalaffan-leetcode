# 3053 — Classifying Triangles By Lengths

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func ClassifyingTrianglesByLengths(triangles [][]int) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3053: Classifying Triangles by Lengths
// https://leetcode.com/problems/classifying-triangles-by-lengths/
// Difficulty: Easy [Paid - Database]
//
// Note: This is a SQL problem on LeetCode. In Go, we implement the
// equivalent logic: classify triangles by side lengths.

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: triangleType
	// Input: [side_a, side_b, side_c]
	fmt.Println(ClassifyingTrianglesByLengths([][]int{{3, 3, 3}, {3, 4, 5}, {3, 3, 5}, {1, 2, 3}}))
	// [Equilateral Scalene Isosceles None]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: triangleType
func ClassifyingTrianglesByLengths(triangles [][]int) []string {
	result := make([]string, len(triangles))
	for i, t := range triangles {
		sides := []int{t[0], t[1], t[2]}
  // Urutkan secara ascending — O(n log n)
		sort.Ints(sides)
		a, b, c := sides[0], sides[1], sides[2]

		if a+b <= c {
			result[i] = "None"
		} else if a == b && b == c {
			result[i] = "Equilateral"
		} else if a == b || b == c || a == c {
			result[i] = "Isosceles"
		} else {
			result[i] = "Scalene"
		}
	}
	return result
}
```
