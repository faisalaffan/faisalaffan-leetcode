# 2881 — Create A New Column

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CreateANewColumn(df [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2881: Create a New Column
// https://leetcode.com/problems/create-a-new-column/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we add a "grade" column computed from existing data.

import "fmt"

func main() {
	// LeetCode name: createBonusColumn
	fmt.Println(CreateANewColumn([][]int{{101, 15}, {102, 11}, {103, 20}}))
	// [[101 15 30] [102 11 22] [103 20 40]]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: createBonusColumn
func CreateANewColumn(df [][]int) [][]int {
  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, len(df))
	for i, row := range df {
		// bonus = salary * 2
		result[i] = []int{row[0], row[1], row[1] * 2}
	}
	return result
}
```
