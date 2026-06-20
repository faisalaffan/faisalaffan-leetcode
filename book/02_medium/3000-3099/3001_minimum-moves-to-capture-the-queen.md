# 3001 — Minimum Moves To Capture The Queen

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minMovesToCaptureTheQueen(a, b, c, d, e, f int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3001: Minimum Moves to Capture The Queen
// https://leetcode.com/problems/minimum-moves-to-capture-the-queen/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(minMovesToCaptureTheQueen(1, 1, 8, 8, 2, 3))
	fmt.Println(minMovesToCaptureTheQueen(5, 3, 3, 4, 5, 2))
}

func minMovesToCaptureTheQueen(a, b, c, d, e, f int) int {
	// Rook and queen on same row, bishop not in between
	if a == e && !(a == c && (d-b)*(d-f) < 0) {
		return 1
	}
	// Rook and queen on same column, bishop not in between
	if b == f && !(b == d && (c-a)*(c-e) < 0) {
		return 1
	}
	// Bishop and queen on same diagonal, rook not in between
	if c+d == e+f && !(a+b == e+f && (a-c)*(a-e) < 0) {
		return 1
	}
	if c-d == e-f && !(a-b == e-f && (a-c)*(a-e) < 0) {
		return 1
	}
	return 2
}
```
