# 1954 — Minimum Garden Perimeter To Collect Enough Apples

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumPerimeter(neededApples int64) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(cuberoot(n)), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1954: Minimum Garden Perimeter to Collect Enough Apples
// https://leetcode.com/problems/minimum-garden-perimeter-to-collect-enough-apples/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinimumPerimeter(1))
	fmt.Println(MinimumPerimeter(13))
	fmt.Println(MinimumPerimeter(1000000000))
}

// Time: O(cuberoot(n)), Space: O(1)
func MinimumPerimeter(neededApples int64) int64 {
	// For a garden with side length 2n (total apples = 2n(n+1)(2n+1))
	// Apples = 2 * n * (n+1) * (2n+1)
	// Perimeter = 8 * n

	n := int64(1)
	for {
		apples := 2 * n * (n + 1) * (2*n + 1)
		if apples >= neededApples {
			return 8 * n
		}
		n++
	}
}
```
