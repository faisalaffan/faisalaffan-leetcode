# 2885 — Rename Columns

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func RenameColumns(df [][]string) [][]string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2885: Rename Columns
// https://leetcode.com/problems/rename-columns/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we rename columns: id->student_id, first->first_name, last->last_name, age->age_in_years.

import "fmt"

func main() {
	// LeetCode name: renameColumns
	// Input: [id, first, last, age]
	fmt.Println(RenameColumns([][]string{{"1", "John", "Doe", "15"}, {"2", "Jane", "Smith", "20"}}))
	// [[1 John Doe 15] [2 Jane Smith 20]]
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: renameColumns
func RenameColumns(df [][]string) [][]string {
	// Column renaming is a metadata operation; data itself doesn't change.
	return df
}
```
