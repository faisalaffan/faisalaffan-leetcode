# 1288 — Remove Covered Intervals

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func removeCoveredIntervals(intervals [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1288: Remove Covered Intervals
// https://leetcode.com/problems/remove-covered-intervals/
// Difficulty: Medium

// Sort by start ascending, end descending. Track current max end.
// If end <= maxEnd, interval is covered.

// Time: O(n log n)
// Space: O(1)

func removeCoveredIntervals(intervals [][]int) int {
  // Custom sort dengan comparator
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] > intervals[j][1]
	})

	count := 0
	maxEnd := 0

	for _, inv := range intervals {
		if inv[1] > maxEnd {
			count++
			maxEnd = inv[1]
		}
	}

	return count
}

func main() {
	fmt.Printf("%d (expected: 2)\n",
		removeCoveredIntervals([][]int{{1, 4}, {3, 6}, {2, 8}}))

	fmt.Printf("%d (expected: 1)\n",
		removeCoveredIntervals([][]int{{1, 4}, {2, 3}}))

	fmt.Printf("%d (expected: 2)\n",
		removeCoveredIntervals([][]int{{0, 10}, {5, 12}}))
}
```
