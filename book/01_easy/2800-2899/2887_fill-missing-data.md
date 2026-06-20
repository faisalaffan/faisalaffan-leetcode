# 2887 — Fill Missing Data

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FillMissingData(df [][]string) [][]string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2887: Fill Missing Data
// https://leetcode.com/problems/fill-missing-data/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we fill missing (zero/empty) quantity values with 0.

import "fmt"

func main() {
	// LeetCode name: fillMissingValues
	// Input: [name, quantity]. Empty string means missing.
	fmt.Println(FillMissingData([][]string{{"A", "10"}, {"B", ""}, {"C", "5"}, {"D", ""}}))
	// [[A 10] [B 0] [C 5] [D 0]]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: fillMissingValues
func FillMissingData(df [][]string) [][]string {
  // Membuat matriks/slice 2D untuk DP
	result := make([][]string, len(df))
	for i, row := range df {
		if row[1] == "" {
			result[i] = []string{row[0], "0"}
		} else {
			result[i] = row
		}
	}
	return result
}
```
