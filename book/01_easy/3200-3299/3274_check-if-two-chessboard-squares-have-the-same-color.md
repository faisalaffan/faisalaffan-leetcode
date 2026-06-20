# 3274 — Check If Two Chessboard Squares Have The Same Color

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckIfTwoChessboardSquaresHaveTheSameColor(coordinate1 string, coordinate2 string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3274: Check if Two Chessboard Squares Have the Same Color
// https://leetcode.com/problems/check-if-two-chessboard-squares-have-the-same-color/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckIfTwoChessboardSquaresHaveTheSameColor("a1", "c3"))
	fmt.Println(CheckIfTwoChessboardSquaresHaveTheSameColor("a1", "h3"))
}

// CheckIfTwoChessboardSquaresHaveTheSameColor returns true if both squares are the same color on a chessboard.
// Time: O(1). Space: O(1).
func CheckIfTwoChessboardSquaresHaveTheSameColor(coordinate1 string, coordinate2 string) bool {
	c1 := (int(coordinate1[0]-'a') + int(coordinate1[1]-'1')) % 2
	c2 := (int(coordinate2[0]-'a') + int(coordinate2[1]-'1')) % 2
	return c1 == c2
}
```
