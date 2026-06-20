# 2882 — Drop Duplicate Rows

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func DropDuplicateRows(df [][]string) [][]string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2882: Drop Duplicate Rows
// https://leetcode.com/problems/drop-duplicate-rows/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we remove rows with duplicate emails.

import "fmt"

func main() {
	// LeetCode name: dropDuplicateEmails
	fmt.Println(DropDuplicateRows([][]string{{"1", "a@b.com"}, {"2", "c@d.com"}, {"3", "a@b.com"}}))
	// [[1 a@b.com] [2 c@d.com]]

	fmt.Println(DropDuplicateRows([][]string{{"1", "x@y.com"}, {"2", "x@y.com"}, {"3", "x@y.com"}}))
	// [[1 x@y.com]]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: dropDuplicateEmails
func DropDuplicateRows(df [][]string) [][]string {
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[string]bool)
	result := [][]string{}
	for _, row := range df {
		email := row[1]
		if !seen[email] {
			seen[email] = true
			result = append(result, row)
		}
	}
	return result
}
```
