# 0050 — Powx N

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func myPow(x float64, n int) float64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #50: Pow(x, n)
// https://leetcode.com/problems/powx-n/
// Difficulty: Medium

import "fmt"

func myPow(x float64, n int) float64 {
  // Edge case: input kosong — langsung return
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}

	result := 1.0
	for n > 0 {
		if n&1 == 1 {
			result *= x
		}
		x *= x
		n >>= 1
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(myPow(2.0, 10)) // 1024

	// Test case 2
	fmt.Println(myPow(2.1, 3)) // 9.261

	// Test case 3
	fmt.Println(myPow(2.0, -2)) // 0.25
}

// Time: O(log n) | Space: O(1)
```
