# 0391 — Perfect Rectangle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func isRectangleCover(rectangles [][]int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Bitmask

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #391: Perfect Rectangle
// https://leetcode.com/problems/perfect-rectangle/
// Difficulty: Hard
//
// A perfect rectangle exactly covers its bounding box with no gaps or
// overlaps. Approach: (1) sum of areas == bounding box area; (2) track
// corner parity — internal corners appear an even number of times, valid
// corners exactly once.

import (
	"fmt"
)

func main() {
	// Example 1: true
	fmt.Println(isRectangleCover([][]int{
		{1, 1, 3, 3}, {3, 1, 4, 2}, {3, 2, 4, 4}, {1, 3, 2, 4}, {2, 3, 3, 4},
	}))
	// Example 2: false (gap)
	fmt.Println(isRectangleCover([][]int{
		{1, 1, 2, 3}, {1, 3, 2, 4}, {3, 1, 4, 2}, {3, 2, 4, 4},
	}))
	// Example 3: false (overlap)
	fmt.Println(isRectangleCover([][]int{
		{1, 1, 3, 3}, {3, 1, 4, 2}, {1, 3, 2, 4}, {3, 2, 4, 4},
	}))
	// Edge: single rectangle
	fmt.Println(isRectangleCover([][]int{{0, 0, 1, 1}}))
	// Edge: two adjacent
	fmt.Println(isRectangleCover([][]int{{0, 0, 1, 1}, {0, 1, 1, 2}}))
}

func isRectangleCover(rectangles [][]int) bool {
	if len(rectangles) == 0 {
		return false
	}

	type point struct{ x, y int }
  // Membuat map (HashMap) — pencarian O(1)
	corners := make(map[point]int)

	minX, minY, maxX, maxY := 1<<30, 1<<30, -1<<30, -1<<30
	totalArea := 0

	for _, r := range rectangles {
		x1, y1, x2, y2 := r[0], r[1], r[2], r[3]
		totalArea += (x2 - x1) * (y2 - y1)

		if x1 < minX {
			minX = x1
		}
		if y1 < minY {
			minY = y1
		}
		if x2 > maxX {
			maxX = x2
		}
		if y2 > maxY {
			maxY = y2
		}

		// Toggle corner parity
		corners[point{x1, y1}]++
		corners[point{x1, y2}]++
		corners[point{x2, y1}]++
		corners[point{x2, y2}]++
	}

	// Area check
	bBoxArea := (maxX - minX) * (maxY - minY)
	if totalArea != bBoxArea {
		return false
	}

	// After all toggles, exactly 4 corners should remain (odd count = 1)
	// All internal corners must have even count.
	expected := map[point]bool{
		{minX, minY}: true,
		{minX, maxY}: true,
		{maxX, minY}: true,
		{maxX, maxY}: true,
	}

	for p, cnt := range corners {
		if cnt%2 != 0 {
			if !expected[p] {
				return false
			}
		}
	}

	// Verify that bounding box corners appear exactly once
	for p := range expected {
		if corners[p]%2 != 1 {
			return false
		}
	}

	return true
}
```
