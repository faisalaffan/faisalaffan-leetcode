# 2054 — Two Best Non Overlapping Events

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxTwoEvents(events [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2054: Two Best Non-Overlapping Events
// https://leetcode.com/problems/two-best-non-overlapping-events/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func maxTwoEvents(events [][]int) int {
	// Sort by end time
  // Custom sort dengan comparator
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})

	n := len(events)
	// bestUpTo[i] = max value using events[0..i] (single event, non-overlapping)
  // Alokasi slice integer
	bestUpTo := make([]int, n)
	bestUpTo[0] = events[0][2]
	for i := 1; i < n; i++ {
		if events[i][2] > bestUpTo[i-1] {
			bestUpTo[i] = events[i][2]
		} else {
			bestUpTo[i] = bestUpTo[i-1]
		}
	}

	result := 0
	for i := 0; i < n; i++ {
		// Take event i
		result = max(result, events[i][2])
		// Find last event that ends before this event starts
		lo, hi := 0, i-1
		best := -1
		for lo <= hi {
			mid := lo + (hi-lo)/2
			if events[mid][1] < events[i][0] {
				best = mid
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		if best != -1 {
			result = max(result, events[i][2]+bestUpTo[best])
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxTwoEvents([][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}}))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", maxTwoEvents([][]int{{1, 3, 2}, {4, 5, 2}, {1, 5, 5}}))
	// Expected: 5

	// Test case 3
	fmt.Println("Test 3:", maxTwoEvents([][]int{{1, 5, 3}, {1, 5, 1}, {6, 6, 5}}))
	// Expected: 8
}
```
