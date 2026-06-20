# 1812 — Determine Color Of A Chessboard Square

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SquareIsWhite(coordinates string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1812: Determine Color of a Chessboard Square
// https://leetcode.com/problems/determine-color-of-a-chessboard-square/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func SquareIsWhite(coordinates string) bool {
	// (col + row) % 2 == 0 means dark, == 1 means light (white)
	col := int(coordinates[0] - 'a' + 1)
	row := int(coordinates[1] - '0')
	return (col+row)%2 == 1
}

func main() {
	fmt.Println(SquareIsWhite("a1"))
	fmt.Println(SquareIsWhite("h3"))
	fmt.Println(SquareIsWhite("c7"))
}
```
