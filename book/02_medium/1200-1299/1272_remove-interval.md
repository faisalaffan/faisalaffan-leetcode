# 1272 — Remove Interval

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func removeInterval(intervals [][]int, toBeRemoved []int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1272: Remove Interval
// https://leetcode.com/problems/remove-interval/
// Difficulty: Medium [Paid]

// Remove an interval from a set of non-overlapping intervals.
// Return the resulting intervals.

// Time: O(n)
// Space: O(n)

func removeInterval(intervals [][]int, toBeRemoved []int) [][]int {
  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, 0)
	rs, re := toBeRemoved[0], toBeRemoved[1]

	for _, interval := range intervals {
		s, e := interval[0], interval[1]
		// No overlap
		if e <= rs || s >= re {
			result = append(result, interval)
		} else {
			// Left part (if any)
			if s < rs {
				result = append(result, []int{s, rs})
			}
			// Right part (if any)
			if e > re {
				result = append(result, []int{re, e})
			}
		}
	}

	return result
}

func main() {
	fmt.Printf("%v (expected: [[0 1] [6 7]])\n",
		removeInterval([][]int{{0, 2}, {3, 4}, {5, 7}}, []int{1, 6}))

	fmt.Printf("%v (expected: [[3 4]])\n",
		removeInterval([][]int{{0, 5}}, []int{0, 3}))
}
```
