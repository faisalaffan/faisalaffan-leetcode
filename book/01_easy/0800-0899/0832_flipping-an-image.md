# 0832 — Flipping An Image

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func flipAndInvertImage(image [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m*n). Space: O(1) in-place.  
**Kompleksitas Ruang:** O(1) in-place.

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #832: Flipping an Image
// https://leetcode.com/problems/flipping-an-image/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}})) // [[1,0,0],[0,1,0],[1,1,1]]
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}))
}

// flipAndInvertImage flips the image horizontally then inverts it.
// Time: O(m*n). Space: O(1) in-place.
func flipAndInvertImage(image [][]int) [][]int {
	for _, row := range image {
		l, r := 0, len(row)-1
		for l <= r {
			row[l], row[r] = 1-row[r], 1-row[l]
			l++
			r--
		}
	}
	return image
}
```
