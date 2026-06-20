# 2886 — Change Data Type

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ChangeDataType(df [][]float64) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2886: Change Data Type
// https://leetcode.com/problems/change-data-type/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we convert the grade column from float64 to int (truncation).

import "fmt"

func main() {
	// LeetCode name: changeDataType
	// Input: [student_id, grade (float)]
	fmt.Println(ChangeDataType([][]float64{{1, 3.5}, {2, 4.2}, {3, 2.8}}))
	// [[1 3] [2 4] [3 2]]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: changeDataType
func ChangeDataType(df [][]float64) [][]int {
  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, len(df))
	for i, row := range df {
		result[i] = []int{int(row[0]), int(row[1])}
	}
	return result
}
```
