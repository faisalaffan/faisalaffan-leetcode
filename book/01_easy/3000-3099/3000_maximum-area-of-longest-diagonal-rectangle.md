# 3000 — Maximum Area Of Longest Diagonal Rectangle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumAreaOfLongestDiagonalRectangle(dimensions [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3000: Maximum Area of Longest Diagonal Rectangle
// https://leetcode.com/problems/maximum-area-of-longest-diagonal-rectangle/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: areaOfMaxDiagonal
	fmt.Println(MaximumAreaOfLongestDiagonalRectangle([][]int{{9, 3}, {8, 6}})) // 48
	fmt.Println(MaximumAreaOfLongestDiagonalRectangle([][]int{{3, 4}, {4, 3}})) // 12
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: areaOfMaxDiagonal
func MaximumAreaOfLongestDiagonalRectangle(dimensions [][]int) int {
	maxDiag := 0
	maxArea := 0
	for _, dim := range dimensions {
		l, w := dim[0], dim[1]
		diagSq := l*l + w*w
		area := l * w
		if diagSq > maxDiag || (diagSq == maxDiag && area > maxArea) {
			maxDiag = diagSq
			maxArea = area
		}
	}
	return maxArea
}
```
