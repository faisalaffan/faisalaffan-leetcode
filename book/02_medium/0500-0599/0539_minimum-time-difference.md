# 0539 — Minimum Time Difference

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindMinDifference(timePoints []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #539: Minimum Time Difference
// https://leetcode.com/problems/minimum-time-difference/
// Difficulty: Medium
// Time: O(n log n)
// Space: O(n)

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(FindMinDifference([]string{"23:59", "00:00"}))
	fmt.Println(FindMinDifference([]string{"00:00", "23:59", "00:00"}))
}

func FindMinDifference(timePoints []string) int {
	n := len(timePoints)
  // Alokasi slice integer
	minutes := make([]int, n)

	for i, t := range timePoints {
		h, _ := strconv.Atoi(t[:2])
		m, _ := strconv.Atoi(t[3:])
		minutes[i] = h*60 + m
	}

  // Urutkan secara ascending — O(n log n)
	sort.Ints(minutes)

	minDiff := 24 * 60 // 1440
	for i := 1; i < n; i++ {
		diff := minutes[i] - minutes[i-1]
		if diff < minDiff {
			minDiff = diff
		}
	}

	// Check circular difference
	circularDiff := 1440 - minutes[n-1] + minutes[0]
	if circularDiff < minDiff {
		minDiff = circularDiff
	}

	return minDiff
}
```
