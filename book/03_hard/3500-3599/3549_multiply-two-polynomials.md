# 3549 — Multiply Two Polynomials

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func multiply(poly1 []int, poly2 []int) []int
```

> **💡 Hint:** Standard polynomial multiplication O(n*m).

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3549: Multiply Two Polynomials
// https://leetcode.com/problems/multiply-two-polynomials/
// Difficulty: Hard [Paid]
//
// Given two polynomials represented as arrays of coefficients (index = power),
// return their product as an array of coefficients.
//
// Approach: Standard polynomial multiplication O(n*m).

import "fmt"

func main() {
	// Example 1
	fmt.Println(multiply([]int{1, 2, 3}, []int{4, 5}))
	// Example 2
	fmt.Println(multiply([]int{1, 1}, []int{1, 1}))
	// Example 3: constant polynomial
	fmt.Println(multiply([]int{2}, []int{3, 4}))
	// Edge: with zeros
	fmt.Println(multiply([]int{0, 1}, []int{1, 0, 1}))
}

func multiply(poly1 []int, poly2 []int) []int {
	if len(poly1) == 0 || len(poly2) == 0 {
		return []int{}
	}

  // Alokasi slice integer
	result := make([]int, len(poly1)+len(poly2)-1)
	for i, c1 := range poly1 {
		for j, c2 := range poly2 {
			result[i+j] += c1 * c2
		}
	}
	return result
}
```
