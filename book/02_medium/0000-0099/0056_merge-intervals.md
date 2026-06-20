# 0056 — Merge Intervals

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func merge(intervals [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Merge Sort

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Merge Sort** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #56: Merge Intervals
// https://leetcode.com/problems/merge-intervals/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

  // Custom sort dengan comparator
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		if intervals[i][0] <= last[1] {
			if intervals[i][1] > last[1] {
				last[1] = intervals[i][1]
			}
		} else {
			result = append(result, intervals[i])
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	// [[1 6] [8 10] [15 18]]

	// Test case 2
	fmt.Println(merge([][]int{{1, 4}, {4, 5}})) // [[1 5]]

	// Test case 3
	fmt.Println(merge([][]int{{1, 4}, {2, 3}})) // [[1 4]]
}

// Time: O(n log n) | Space: O(n)
```
