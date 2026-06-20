# 0625 — Minimum Factorization

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SmallestFactorization(num int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #625: Minimum Factorization
// https://leetcode.com/problems/minimum-factorization/
// Difficulty: Medium [Paid]
// Time: O(log n)
// Space: O(1)

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(SmallestFactorization(48))
	fmt.Println(SmallestFactorization(15))
	fmt.Println(SmallestFactorization(1))
}

func SmallestFactorization(num int) int {
	if num < 2 {
		return num
	}

	// Build number from right to left using digits 9..2
	result := 0
	multiplier := 1

	for i := 9; i >= 2; i-- {
		for num%i == 0 {
			result += i * multiplier
			if result > math.MaxInt32 {
				return 0
			}
			multiplier *= 10
			num /= i
		}
	}

	if num > 1 {
		return 0
	}
	return result
}
```
