# 0836 — Rectangle Overlap

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func isRectangleOverlap(rec1 []int, rec2 []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #836: Rectangle Overlap
// https://leetcode.com/problems/rectangle-overlap/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(isRectangleOverlap([]int{0, 0, 2, 2}, []int{1, 1, 3, 3})) // true
	fmt.Println(isRectangleOverlap([]int{0, 0, 1, 1}, []int{1, 0, 2, 1})) // false
	fmt.Println(isRectangleOverlap([]int{0, 0, 1, 1}, []int{2, 2, 3, 3})) // false
}

// isRectangleOverlap checks if two rectangles overlap (positive area).
// Time: O(1). Space: O(1).
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// Check if one rectangle is to the left of the other
	if rec1[2] <= rec2[0] || rec2[2] <= rec1[0] {
		return false
	}
	// Check if one rectangle is above the other
	if rec1[3] <= rec2[1] || rec2[3] <= rec1[1] {
		return false
	}
	return true
}
```
