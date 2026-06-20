# 3899 — Angles Of A Triangle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func AnglesOfATriangle(sides []int) []float64
```

> **💡 Hint:** Use law of cosines. Check triangle inequality first.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
