# 2250 — Count Number Of Rectangles Containing Each Point

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countRectangles(rectangles [][]int, points [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** O((n + m) log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2250: Count Number of Rectangles Containing Each Point
// https://leetcode.com/problems/count-number-of-rectangles-containing-each-point/
// Difficulty: Medium
// Time: O((n + m) log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func countRectangles(rectangles [][]int, points [][]int) []int {
	// Group rectangles by height
  // Membuat matriks/slice 2D untuk DP
	byHeight := make([][]int, 101)
	for _, r := range rectangles {
		h := r[1]
		byHeight[h] = append(byHeight[h], r[0])
	}
	for h := 0; h <= 100; h++ {
  // Urutkan secara ascending — O(n log n)
		sort.Ints(byHeight[h])
	}

  // Alokasi slice integer
	result := make([]int, len(points))
	for i, p := range points {
		x, y := p[0], p[1]
		count := 0
		for h := y; h <= 100; h++ {
			if len(byHeight[h]) == 0 {
				continue
			}
			// Binary search for first rectangle with width >= x
			idx := sort.SearchInts(byHeight[h], x)
			count += len(byHeight[h]) - idx
		}
		result[i] = count
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(countRectangles([][]int{{1, 2}, {2, 3}, {2, 5}}, [][]int{{2, 1}, {1, 4}}))
	// Expected: [2, 1]

	// Test case 2
	fmt.Println(countRectangles([][]int{{1, 1}, {2, 2}, {3, 3}}, [][]int{{1, 3}, {1, 1}}))
	// Expected: [1, 3]
}
```
