# 3414 — Maximum Score Of Non Overlapping Intervals

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumScoreOfNonOverlappingIntervals(intervals [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search, Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3414: Maximum Score of Non-overlapping Intervals
// https://leetcode.com/problems/maximum-score-of-non-overlapping-intervals/
// Difficulty: Hard
//
// Weighted interval scheduling. Sort by end time, binary search for
// previous non-overlapping interval, DP.

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaximumScoreOfNonOverlappingIntervals([][]int{{1, 3, 2}, {2, 5, 3}, {4, 6, 1}}))
}

func MaximumScoreOfNonOverlappingIntervals(intervals [][]int) int {
	n := len(intervals)
	type interval struct{ start, end, score int }
  // Alokasi slice integer
	ivs := make([]interval, n)
	for i, v := range intervals {
		ivs[i] = interval{v[0], v[1], v[2]}
	}
  // Custom sort dengan comparator
	sort.Slice(ivs, func(i, j int) bool {
		return ivs[i].end < ivs[j].end
	})

  // Alokasi slice integer
	ends := make([]int, n)
	for i, v := range ivs {
		ends[i] = v.end
	}

  // Alokasi slice integer
	dp := make([]int, n)
	dp[0] = ivs[0].score
	for i := 1; i < n; i++ {
		// Binary search for last interval ending <= ivs[i].start
		prev := sort.SearchInts(ends, ivs[i].start+1) - 1
		best := ivs[i].score
		if prev >= 0 {
			best += dp[prev]
		}
		if dp[i-1] > best {
			best = dp[i-1]
		}
		dp[i] = best
	}
	return dp[n-1]
}
```
