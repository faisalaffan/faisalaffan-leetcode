# 0172 — Factorial Trailing Zeroes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func trailingZeroes(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #172: Factorial Trailing Zeroes
// https://leetcode.com/problems/factorial-trailing-zeroes/
// Difficulty: Medium
// Time: O(log n), Space: O(1)

import "fmt"

func trailingZeroes(n int) int {
	count := 0
	for n >= 5 {
		n /= 5
		count += n
	}
	return count
}

func main() {
	fmt.Println(trailingZeroes(3))
	fmt.Println(trailingZeroes(5))
	fmt.Println(trailingZeroes(0))
}
```
