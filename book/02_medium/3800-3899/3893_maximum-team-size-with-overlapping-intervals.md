# 3893 — Maximum Team Size With Overlapping Intervals

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumTeamSizeWithOverlappingIntervals(startTime []int, endTime []int) int
```

> **💡 Hint:** For each employee, count overlapping intervals using binary search

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** O(N log N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3893: Maximum Team Size with Overlapping Intervals
// https://leetcode.com/problems/maximum-team-size-with-overlapping-intervals/
// Difficulty: Medium [Paid]
// Time: O(N log N) | Space: O(N)
// Approach: For each employee, count overlapping intervals using binary search
// on sorted start and end times.

import (
	"fmt"
	"sort"
)

func MaximumTeamSizeWithOverlappingIntervals(startTime []int, endTime []int) int {
	n := len(startTime)
  // Alokasi slice integer
	st := make([]int, n)
  // Alokasi slice integer
	et := make([]int, n)
	copy(st, startTime)
	copy(et, endTime)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(st)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(et)

	ans := 0
	for i := 0; i < n; i++ {
		start := startTime[i]
		end := endTime[i]

		// Count employees whose start <= end
		startsBeforeEnd := sort.SearchInts(st, end+1)
		// Count employees whose end < start
		endsBeforeStart := sort.SearchInts(et, start)

		overlap := startsBeforeEnd - endsBeforeStart
		if overlap > ans {
			ans = overlap
		}
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(MaximumTeamSizeWithOverlappingIntervals([]int{1, 2, 3}, []int{4, 5, 6})) // Expected: 3

	// Example 2
	fmt.Println(MaximumTeamSizeWithOverlappingIntervals([]int{2, 5, 8}, []int{3, 7, 9})) // Expected: 1

	// Example 3
	fmt.Println(MaximumTeamSizeWithOverlappingIntervals([]int{3, 4, 6}, []int{8, 5, 7})) // Expected: 3
}
```
