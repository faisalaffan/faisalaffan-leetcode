# 2500 — Delete Greatest Value In Each Row

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func DeleteGreatestValueInEachRow(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2500: Delete Greatest Value in Each Row
// https://leetcode.com/problems/delelete-greatest-value-in-each-row/
// Difficulty: Easy
// Time O(n * m log m) | Space O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(DeleteGreatestValueInEachRow([][]int{{1, 2, 4}, {3, 3, 1}})) // 8
	fmt.Println(DeleteGreatestValueInEachRow([][]int{{10}}))                 // 10
}

func DeleteGreatestValueInEachRow(grid [][]int) int {
  // Range loop: iterasi dengan indeks + nilai
	for i := range grid {
  // Urutkan secara ascending — O(n log n)
		sort.Ints(grid[i])
	}

	m := len(grid[0])
	sum := 0
	for col := m - 1; col >= 0; col-- {
		maxVal := 0
		for row := 0; row < len(grid); row++ {
			if grid[row][col] > maxVal {
				maxVal = grid[row][col]
			}
		}
		sum += maxVal
	}
	return sum
}
```
