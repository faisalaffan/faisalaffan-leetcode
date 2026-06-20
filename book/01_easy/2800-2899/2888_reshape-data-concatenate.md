# 2888 — Reshape Data Concatenate

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ReshapeDataConcatenate(df1, df2 [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n+m)  
**Kompleksitas Ruang:** O(n+m)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2888: Reshape Data: Concatenate
// https://leetcode.com/problems/reshape-data-concatenate/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we concatenate two dataframes vertically.

import "fmt"

func main() {
	// LeetCode name: concatenateDataFrames
	fmt.Println(ReshapeDataConcatenate([][]int{{1, 15}, {2, 11}}, [][]int{{3, 12}, {4, 14}}))
	// [[1 15] [2 11] [3 12] [4 14]]
}

// Time: O(n+m) | Space: O(n+m)
// LeetCode submission name: concatenateDataFrames
func ReshapeDataConcatenate(df1, df2 [][]int) [][]int {
  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, 0, len(df1)+len(df2))
	result = append(result, df1...)
	result = append(result, df2...)
	return result
}
```
