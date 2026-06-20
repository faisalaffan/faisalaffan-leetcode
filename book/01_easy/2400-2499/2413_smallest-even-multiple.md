# 2413 — Smallest Even Multiple

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SmallestEvenMultiple(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2413: Smallest Even Multiple
// https://leetcode.com/problems/smallest-even-multiple/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(SmallestEvenMultiple(5)) // 10
	fmt.Println(SmallestEvenMultiple(6)) // 6
}

func SmallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	}
	return n * 2
}
```
