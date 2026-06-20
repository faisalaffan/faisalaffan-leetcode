# 2280 — Minimum Lines To Represent A Line Chart

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumLines(stockPrices [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2280: Minimum Lines to Represent a Line Chart
// https://leetcode.com/problems/minimum-lines-to-represent-a-line-chart/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func minimumLines(stockPrices [][]int) int {
	if len(stockPrices) <= 1 {
		return 0
	}

  // Custom sort dengan comparator
	sort.Slice(stockPrices, func(i, j int) bool {
		return stockPrices[i][0] < stockPrices[j][0]
	})

	lines := 1
	for i := 2; i < len(stockPrices); i++ {
		x1, y1 := stockPrices[i-2][0], stockPrices[i-2][1]
		x2, y2 := stockPrices[i-1][0], stockPrices[i-1][1]
		x3, y3 := stockPrices[i][0], stockPrices[i][1]

		// Compare slopes: (y2-y1)/(x2-x1) == (y3-y2)/(x3-x2)
		// Cross multiply to avoid floating point
		if (y2-y1)*(x3-x2) != (y3-y2)*(x2-x1) {
			lines++
		}
	}
	return lines
}

func main() {
	// Test case 1
	fmt.Println(minimumLines([][]int{{1, 7}, {2, 6}, {3, 5}, {4, 4}, {5, 4}, {6, 3}, {7, 2}, {8, 1}}))
	// Expected: 3

	// Test case 2
	fmt.Println(minimumLines([][]int{{3, 4}, {1, 2}, {7, 8}, {2, 3}}))
	// Expected: 1

	// Test case 3
	fmt.Println(minimumLines([][]int{{1, 1}}))
	// Expected: 0
}
```
