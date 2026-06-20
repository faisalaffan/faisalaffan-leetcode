# 2194 — Cells In A Range On An Excel Sheet

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CellsInARangeOnAnExcelSheet(s string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O((colDiff+1) * (rowDiff+1)), Space: O((colDiff+1) * (rowDiff+1))  
**Kompleksitas Ruang:** O((colDiff+1) * (rowDiff+1))

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2194: Cells in a Range on an Excel Sheet
// https://leetcode.com/problems/cells-in-a-range-on-an-excel-sheet/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CellsInARangeOnAnExcelSheet("K1:L2")) // [K1 K2 L1 L2]
	fmt.Println(CellsInARangeOnAnExcelSheet("A1:F1")) // [A1 B1 C1 D1 E1 F1]
}

// Time: O((colDiff+1) * (rowDiff+1)), Space: O((colDiff+1) * (rowDiff+1))
func CellsInARangeOnAnExcelSheet(s string) []string {
	col1 := s[0]
	row1 := s[1]
	col2 := s[3]
	row2 := s[4]

	var result []string
	for c := col1; c <= col2; c++ {
		for r := row1; r <= row2; r++ {
			result = append(result, string(c)+string(r))
		}
	}
	return result
}
```
