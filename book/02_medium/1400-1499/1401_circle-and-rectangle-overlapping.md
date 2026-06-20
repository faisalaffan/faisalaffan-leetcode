# 1401 — Circle And Rectangle Overlapping

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1401: Circle and Rectangle Overlapping
// https://leetcode.com/problems/circle-and-rectangle-overlapping/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(checkOverlap(1, 0, 0, 1, -1, 3, 1)) // true

	// Test case 2
	fmt.Println(checkOverlap(1, 1, 1, -3, -3, 3, 3)) // true

	// Test case 3
	fmt.Println(checkOverlap(1, 0, 0, -1, 0, 0, 1)) // false

	// Test case 4
	fmt.Println(checkOverlap(1, 1, 1, -3, -3, 3, -1)) // true
}

// Time: O(1)
// Space: O(1)
func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
	// Find the closest point on the rectangle to the circle center
	closestX := max(x1, min(x2, xCenter))
	closestY := max(y1, min(y2, yCenter))

	// Calculate distance from circle center to closest point
	dx := xCenter - closestX
	dy := yCenter - closestY

	return dx*dx+dy*dy <= radius*radius
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
