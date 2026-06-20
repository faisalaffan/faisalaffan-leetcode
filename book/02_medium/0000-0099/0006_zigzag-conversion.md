# 0006 — Zigzag Conversion

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func convert(s string, numRows int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #6: Zigzag Conversion
// https://leetcode.com/problems/zigzag-conversion/
// Difficulty: Medium

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

  // Membuat matriks/slice 2D untuk DP
	rows := make([][]byte, numRows)
	curRow := 0
	goingDown := false

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		rows[curRow] = append(rows[curRow], s[i])
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}

	result := make([]byte, 0, len(s))
	for _, row := range rows {
		result = append(result, row...)
	}

	return string(result)
}

func main() {
	// Test case 1
	fmt.Println(convert("PAYPALISHIRING", 3)) // "PAHNAPLSIIGYIR"

	// Test case 2
	fmt.Println(convert("PAYPALISHIRING", 4)) // "PINALSIGYAHRPI"

	// Test case 3
	fmt.Println(convert("A", 1)) // "A"
}

// Time: O(n) | Space: O(n)
```
