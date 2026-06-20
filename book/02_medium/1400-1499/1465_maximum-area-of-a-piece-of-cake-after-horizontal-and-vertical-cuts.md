# 1465 — Maximum Area Of A Piece Of Cake After Horizontal And Vertical Cuts

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n + m log m) for sorting  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1465: Maximum Area of a Piece of Cake After Horizontal and Vertical Cuts
// https://leetcode.com/problems/maximum-area-of-a-piece-of-cake-after-horizontal-and-vertical-cuts/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// Test case 1
	fmt.Println(maxArea(5, 4, []int{1, 2, 4}, []int{1, 3})) // 4

	// Test case 2
	fmt.Println(maxArea(5, 4, []int{3, 1}, []int{1})) // 6

	// Test case 3
	fmt.Println(maxArea(5, 4, []int{3}, []int{3})) // 9
}

const mod = 1_000_000_007

// Time: O(n log n + m log m) for sorting
// Space: O(1)
func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(horizontalCuts)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(verticalCuts)

	// Find max gap in horizontal cuts (including edges)
	maxHDiff := max(horizontalCuts[0], h-horizontalCuts[len(horizontalCuts)-1])
	for i := 1; i < len(horizontalCuts); i++ {
		diff := horizontalCuts[i] - horizontalCuts[i-1]
		if diff > maxHDiff {
			maxHDiff = diff
		}
	}

	// Find max gap in vertical cuts (including edges)
	maxVDiff := max(verticalCuts[0], w-verticalCuts[len(verticalCuts)-1])
	for i := 1; i < len(verticalCuts); i++ {
		diff := verticalCuts[i] - verticalCuts[i-1]
		if diff > maxVDiff {
			maxVDiff = diff
		}
	}

	area := (maxHDiff % mod) * (maxVDiff % mod) % mod
	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
