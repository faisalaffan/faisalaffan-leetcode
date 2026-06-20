# 1037 — Valid Boomerang

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func isBoomerang(points [][]int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1037: Valid Boomerang
// https://leetcode.com/problems/valid-boomerang/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(isBoomerang([][]int{{1, 1}, {2, 3}, {3, 2}})) // true
	fmt.Println(isBoomerang([][]int{{1, 1}, {2, 2}, {3, 3}})) // false
	fmt.Println(isBoomerang([][]int{{0, 0}, {0, 2}, {0, 1}})) // false
}

// LeetCode submission: isBoomerang
func isBoomerang(points [][]int) bool {
	x1, y1 := points[0][0], points[0][1]
	x2, y2 := points[1][0], points[1][1]
	x3, y3 := points[2][0], points[2][1]
	// Check area of triangle: (x2-x1)*(y3-y1) != (x3-x1)*(y2-y1)
	return (x2-x1)*(y3-y1) != (x3-x1)*(y2-y1)
}
```
