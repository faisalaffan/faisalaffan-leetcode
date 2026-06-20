# 1943 — Describe The Painting

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SplitPainting(segments [][]int) [][]int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1943: Describe the Painting
// https://leetcode.com/problems/describe-the-painting/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SplitPainting([][]int{{1, 4, 5}, {4, 7, 7}, {1, 7, 9}}))
	fmt.Println(SplitPainting([][]int{{1, 7, 9}, {6, 8, 15}, {8, 10, 7}}))
	fmt.Println(SplitPainting([][]int{{1, 4, 5}, {1, 4, 7}, {4, 7, 1}, {4, 7, 11}}))
}

// Time: O(n log n), Space: O(n)
func SplitPainting(segments [][]int) [][]int64 {
  // Membuat map (HashMap) — pencarian O(1)
	diff := make(map[int]int64)
  // Membuat map (HashMap) — pencarian O(1)
	endpoints := make(map[int]bool)

	for _, seg := range segments {
		start, end, color := seg[0], seg[1], seg[2]
		diff[start] += int64(color)
		diff[end] -= int64(color)
		endpoints[start] = true
		endpoints[end] = true
	}

	// Sort unique endpoints
  // Alokasi slice integer
	points := make([]int, 0, len(endpoints))
	for p := range endpoints {
		points = append(points, p)
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(points)

  // Membuat matriks/slice 2D untuk DP
	result := make([][]int64, 0)
	var sum int64 = 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(points)-1; i++ {
		sum += diff[points[i]]
		if sum != 0 {
			result = append(result, []int64{int64(points[i]), int64(points[i+1]), sum})
		}
	}
	return result
}
```
