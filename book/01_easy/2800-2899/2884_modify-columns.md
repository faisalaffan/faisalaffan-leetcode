# 2884 — Modify Columns

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ModifyColumns(df [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2884: Modify Columns
// https://leetcode.com/problems/modify-columns/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we modify a column by multiplying salary by 2.

import "fmt"

func main() {
	// LeetCode name: modifySalaryColumn
	fmt.Println(ModifyColumns([][]int{{1, 100}, {2, 200}, {3, 300}}))
	// [[1 200] [2 400] [3 600]]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: modifySalaryColumn
func ModifyColumns(df [][]int) [][]int {
  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, len(df))
	for i, row := range df {
		result[i] = []int{row[0], row[1] * 2}
	}
	return result
}
```
