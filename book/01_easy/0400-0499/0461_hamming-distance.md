# 0461 — Hamming Distance

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func HammingDistance(x, y int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #461: Hamming Distance
// https://leetcode.com/problems/hamming-distance/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func HammingDistance(x, y int) int {
	xor := x ^ y
	count := 0
	for xor > 0 {
		xor &= xor - 1
		count++
	}
	return count
}

func main() {
	fmt.Println(HammingDistance(1, 4))
	fmt.Println(HammingDistance(3, 1))
}
```
