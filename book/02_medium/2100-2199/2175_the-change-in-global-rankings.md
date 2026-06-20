# 2175 — The Change In Global Rankings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func globalRankings(pointsBefore []int, pointsAfter []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2175: The Change in Global Rankings
// https://leetcode.com/problems/the-change-in-global-rankings/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func globalRankings(pointsBefore []int, pointsAfter []int) []int {
	n := len(pointsBefore)

	// Create sorted list of (points, originalIndex) for before
	type pair struct {
		points int
		idx    int
	}
	before := make([]pair, n)
	for i := 0; i < n; i++ {
		before[i] = pair{pointsBefore[i], i}
	}
  // Custom sort dengan comparator
	sort.Slice(before, func(i, j int) bool {
		if before[i].points != before[j].points {
			return before[i].points > before[j].points
		}
		return before[i].idx < before[j].idx
	})

	// Compute rank before: rank = position (1-indexed) when sorted descending
  // Alokasi slice integer
	rankBefore := make([]int, n)
	for pos, p := range before {
		rankBefore[p.idx] = pos + 1
	}

	// Create sorted list for after
	after := make([]pair, n)
	for i := 0; i < n; i++ {
		after[i] = pair{pointsAfter[i], i}
	}
  // Custom sort dengan comparator
	sort.Slice(after, func(i, j int) bool {
		if after[i].points != after[j].points {
			return after[i].points > after[j].points
		}
		return after[i].idx < after[j].idx
	})

	// Compute rank after
  // Alokasi slice integer
	rankAfter := make([]int, n)
	for pos, p := range after {
		rankAfter[p.idx] = pos + 1
	}

	// Difference
  // Alokasi slice integer
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = rankBefore[i] - rankAfter[i]
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", globalRankings([]int{10, 20, 30}, []int{15, 25, 35}))
	// Expected: [0, 0, 0] (same relative order)

	// Test case 2
	fmt.Println("Test 2:", globalRankings([]int{50, 40, 30, 20}, []int{45, 45, 35, 25}))
	// Expected: varying based on rank changes
}
```
