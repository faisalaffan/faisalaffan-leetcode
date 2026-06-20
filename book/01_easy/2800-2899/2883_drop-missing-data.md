# 2883 — Drop Missing Data

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func DropMissingData(df [][]string) [][]string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2883: Drop Missing Data
// https://leetcode.com/problems/drop-missing-data/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we remove rows where the name column is empty.

import "fmt"

func main() {
	// LeetCode name: dropMissingData
	fmt.Println(DropMissingData([][]string{{"1", "Alice", "15"}, {"2", "", "11"}, {"3", "Bob", "12"}}))
	// [[1 Alice 15] [3 Bob 12]]

	fmt.Println(DropMissingData([][]string{{"1", "", "10"}}))
	// []
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: dropMissingData
func DropMissingData(df [][]string) [][]string {
	result := [][]string{}
	for _, row := range df {
		if row[1] != "" {
			result = append(result, row)
		}
	}
	return result
}
```
