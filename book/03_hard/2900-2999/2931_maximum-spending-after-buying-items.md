# 2931 — Maximum Spending After Buying Items

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxSpending(values [][]int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2931: Maximum Spending After Buying Items
// https://leetcode.com/problems/maximum-spending-after-buying-items/
// Difficulty: Hard
//
// m x n grid where each row is sorted ascending. On day d (1-indexed), buy one
// item from the first remaining element of any row. The spending on day d is
// d * (item value). Maximize total spending.
//
// By the rearrangement inequality, spending is maximized by buying items in
// ascending order of value (smallest first, largest last). Since each row is
// sorted ascending, we can simply flatten all values into one sorted list.
// O(m*n log(m*n)) time, O(m*n) space.

import (
	"fmt"
	"sort"
)

func maxSpending(values [][]int) int64 {
	m := len(values)
	if m == 0 {
		return 0
	}
	n := len(values[0])

	// Flatten all values into one slice
  // Alokasi slice integer
	flat := make([]int, 0, m*n)
	for _, row := range values {
		flat = append(flat, row...)
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(flat)

	var ans int64
	for day, val := range flat {
		ans += int64(day+1) * int64(val)
	}
	return ans
}

func main() {
	// Example: values=[[8,5,2],[6,4,1],[9,7,3]] => 285
	// Sorted: [1,2,3,4,5,6,7,8,9]
	// Spending: 1*1 + 2*2 + 3*3 + 4*4 + 5*5 + 6*6 + 7*7 + 8*8 + 9*9 = 285
	fmt.Println(maxSpending([][]int{{8, 5, 2}, {6, 4, 1}, {9, 7, 3}}))

	// Single row
	fmt.Println(maxSpending([][]int{{1, 2, 3}}))

	// Simple case
	fmt.Println(maxSpending([][]int{{10, 20}, {5, 15}}))

	// All same values
	fmt.Println(maxSpending([][]int{{5, 5}, {5, 5}}))

	// Single cell
	fmt.Println(maxSpending([][]int{{7}}))

	// One row ascending
	fmt.Println(maxSpending([][]int{{1, 2, 3, 4, 5}}))

	// Empty grid
	fmt.Println(maxSpending([][]int{}))
}
```
