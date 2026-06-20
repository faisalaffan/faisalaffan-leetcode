# 3899 — Angles Of A Triangle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func AnglesOfATriangle(sides []int) []float64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3899: Angles of a Triangle
// https://leetcode.com/problems/angles-of-a-triangle/
// Difficulty: Medium
// Time: O(1) | Space: O(1)
// Approach: Use law of cosines. Check triangle inequality first.

import (
	"fmt"
	"math"
	"sort"
)

func AnglesOfATriangle(sides []int) []float64 {
	a := float64(sides[0])
	b := float64(sides[1])
	c := float64(sides[2])
	s := []float64{a, b, c}
	sort.Float64s(s)
	a, b, c = s[0], s[1], s[2]

	// Triangle inequality
	if a+b <= c {
		return []float64{}
	}

	// Law of cosines: cos(A) = (b^2 + c^2 - a^2) / (2*b*c)
	angleA := math.Acos((b*b + c*c - a*a) / (2 * b * c)) * 180 / math.Pi
	angleB := math.Acos((a*a + c*c - b*b) / (2 * a * c)) * 180 / math.Pi
	angleC := 180 - angleA - angleB

	angles := []float64{math.Round(angleA*1e5) / 1e5, math.Round(angleB*1e5) / 1e5, math.Round(angleC*1e5) / 1e5}
	sort.Float64s(angles)
	return angles
}

func main() {
	// Example 1
	fmt.Println(AnglesOfATriangle([]int{3, 4, 5})) // Expected: [36.86990 53.13010 90.00000]

	// Example 2
	fmt.Println(AnglesOfATriangle([]int{2, 4, 2})) // Expected: []

	// Extra
	fmt.Println(AnglesOfATriangle([]int{1, 1, 1})) // Expected: ~[60 60 60]
}
```
