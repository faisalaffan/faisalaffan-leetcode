# 2159 — Order Two Columns Independently

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func orderColumns(rows [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2159: Order Two Columns Independently
// https://leetcode.com/problems/order-two-columns-independently/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type Row struct {
	col1, col2 int
}

func orderColumns(rows [][]int) [][]int {
	// Sort by col1 ascending, col2 ascending
  // Custom sort dengan comparator
	sort.Slice(rows, func(i, j int) bool {
		if rows[i][0] != rows[j][0] {
			return rows[i][0] < rows[j][0]
		}
		return rows[i][1] < rows[j][1]
	})

	return rows
}

func main() {
	// Test case 1
	data1 := [][]int{{3, 1}, {1, 3}, {2, 2}}
	fmt.Println("Test 1:", orderColumns(data1))
	// Expected: [[1,3],[2,2],[3,1]]

	// Test case 2
	data2 := [][]int{{5, 5}, {1, 1}, {3, 3}}
	fmt.Println("Test 2:", orderColumns(data2))
	// Expected: [[1,1],[3,3],[5,5]]
}
```
