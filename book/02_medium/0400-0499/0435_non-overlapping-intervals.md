# 0435 — Non Overlapping Intervals

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func eraseOverlapIntervals(intervals [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #435: Non-overlapping Intervals
// https://leetcode.com/problems/non-overlapping-intervals/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// Sort by end
  // Custom sort dengan comparator
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 0
	end := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < end {
			count++
		} else {
			end = intervals[i][1]
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
	// Expected: 1

	// Test case 2
	fmt.Println("Test 2:", eraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}}))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", eraseOverlapIntervals([][]int{{1, 2}, {2, 3}}))
	// Expected: 0
}
```
